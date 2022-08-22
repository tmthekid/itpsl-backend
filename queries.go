package main

const CREATE_USER_QUERY = `INSERT INTO users (name, profession, gender, age, final_growth) VALUES ($1, $2, $3, $4, $5) RETURNING id`
const CREATE_INITIAL_INFORMATION = `INSERT INTO initial_informations (city_value, tax_rate, bonds_amount, bonds_interest_rate, fs_id) VALUES ($1, $2, $3, $4, $5)`
const CREATE_PROBABILITIES_QUERY = `INSERT INTO probabilities (condition_id, condition, amount, percentage, fs_id) VALUES ($1, $2, $3, $4, $5)`
const CREATE_EXPENDITURE_ALLOCATIONS_QUERY = `INSERT INTO expenditure_allocations (expenditure_id, expenditure, amount, fs_id) VALUES ($1, $2, $3, $4)`
const CREATE_FINAL_STATS_QUERY = `INSERT INTO final_stats (year, growth, score_id, score_category, user_id) VALUES($1, $2, $3, $4, $5) RETURNING id`
const CREATE_RESULT_VIEW_QUERY = `CREATE VIEW results AS
SELECT
    u.id,
    u.name,
    u.profession,
    u.gender,
    u.age,
    fs."year",
    fs.growth,
    fs.score_category,
    iinf.city_value,
    iinf.tax_rate,
    iinf.bonds_amount,
    iinf.bonds_interest_rate,
    p_ct."drought",
    p_ct."flood",
    p_ct."landslide",
    p_ct."tsunami",
    p_ct."social_unrest",
    p_ct."normal",
    ea_ct."physical_development",
    ea_ct."eco_system_preservation_&_protection",
    ea_ct."economic_development",
    ea_ct."public_utilities_&_transport",
    ea_ct."waste_&_pollution_management",
    ea_ct."education_&_healthcare",
    ea_ct."cultural_&_heritage_management",
    ea_ct."social_support",
    ea_ct."disaster_mitigation_&_management",
    ea_ct."research_innovation_&_development",
    ea_ct."monitoring_&_management",
    u.final_growth
FROM
    users u
    LEFT JOIN final_stats fs ON u.id = fs.user_id
    LEFT JOIN (
        SELECT
            initial_informations.city_value,
            initial_informations.tax_rate,
            initial_informations.bonds_amount,
            initial_informations.bonds_interest_rate,
            initial_informations.fs_id
        FROM
            initial_informations
    ) AS iinf ON fs.id = iinf.fs_id
    LEFT JOIN (
        SELECT
            *
        FROM
            crosstab(
                'SELECT p.fs_id, p.condition, p.percentage FROM probabilities p GROUP BY p.fs_id, p.id ORDER BY 1'
            ) AS p_ct(
                fs_id BIGINT,
                "drought" FLOAT,
                "flood" FLOAT,
                "landslide" FLOAT,
                "pandemic" FLOAT,
                "tsunami" FLOAT,
                "social_unrest" FLOAT,
                "normal" FLOAT
            )
    ) AS p_ct ON fs.id = p_ct.fs_id
    LEFT JOIN (
        SELECT
            *
        FROM
            crosstab(
                'SELECT
fs.id AS fs_id, ea.expenditure, ea.amount
FROM
final_stats fs
LEFT JOIN expenditure_allocations ea ON ea.fs_id = fs.id
GROUP BY fs.id, ea.id
ORDER BY 1'
            ) AS ea_ct (
                id BIGINT,
                "physical_development" FLOAT,
                "eco_system_preservation_&_protection" FLOAT,
                "economic_development" FLOAT,
                "public_utilities_&_transport" FLOAT,
                "waste_&_pollution_management" FLOAT,
                "education_&_healthcare" FLOAT,
                "cultural_&_heritage_management" FLOAT,
                "social_support" FLOAT,
                "disaster_mitigation_&_management" FLOAT,
                "research_innovation_&_development" FLOAT,
                "monitoring_&_management" FLOAT
            )
    ) AS ea_ct ON fs.id = ea_ct.id`
const GET_CONDITIONS_QUERY = `SELECT * FROM conditions ORDER BY id ASC`
const GET_EXPENDITURE_QUERY = `SELECT * FROM expenditures ORDER BY id ASC`
const GET_SCORES_QUERY = `SELECT * FROM scores ORDER BY id ASC`
const GET_RESULT_QUERY = `SELECT * FROM results ORDER BY id ASC`
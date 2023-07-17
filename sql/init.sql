DROP TABLE IF EXISTS vegetable;

CREATE TABLE vegetable(
    id SERIAL PRIMARY KEY,
    ref_date VARCHAR(255),
    geo VARCHAR(255),
    dguid VARCHAR(255),
    type_of_product VARCHAR(255),
    type_of_storage VARCHAR(255),
    uom VARCHAR(255),
    uom_id VARCHAR(255),
    scalar_factor VARCHAR(255),
    scalar_id VARCHAR(255),
    vector VARCHAR(255),
    coordinate VARCHAR(255),
    value VARCHAR(255),
    status VARCHAR(255),
    symbol VARCHAR(255),
    terminated VARCHAR(255),
    decimals VARCHAR(255)
);

COPY vegetable(ref_date, geo, dguid, type_of_product, type_of_storage,
    uom, uom_id, scalar_factor, scalar_id, vector, coordinate,
    value, status, symbol, terminated, decimals)
    FROM 'C:\Users\Public\32100260.csv'
    DELIMITER ','
    CSV HEADER;
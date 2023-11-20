CREATE TABLE IF NOT EXISTS Api (
    api_id INT NOT NULL AUTO_INCREMENT,
    api_name VARCHAR(50) NOT NULL DEFAULT '',
    api_description VARCHAR(255) NOT NULL DEFAULT '',
    PRIMARY KEY (api_id)
);

CREATE TABLE IF NOT EXISTS ApiSetting (
    api_setting_id INT NOT NULL AUTO_INCREMENT,
    api_id INT NOT NULL,
    reqiuest_method VARCHAR(10) NOT NULL DEFAULT 'POST',
    endpoint VARCHAR(50) NOT NULL DEFAULT '',
    execution_interval_sec INT NOT NULL DEFAULT 0,
    PRIMARY KEY (api_setting_id),
    FOREIGN KEY (api_id) REFERENCES Api (api_id)
);

CREATE TABLE IF NOT EXISTS ApiHeaderSetting (
    api_setting_id INT NOT NULL,
    api_header_setting_no INT NOT NULL,
    api_header_key VARCHAR(50) NOT NULL DEFAULT '',
    api_header_value VARCHAR(50) NOT NULL DEFAULT '',
    PRIMARY KEY (api_setting_id, api_header_setting_no),
    FOREIGN KEY (api_setting_id) REFERENCES ApiSetting (api_setting_id)
);

CREATE TABLE IF NOT EXISTS ApiParamSetting (
    api_setting_id INT NOT NULL,
    api_param_setting_no INT NOT NULL,
    api_param_key VARCHAR(50) NOT NULL DEFAULT '',
    api_param_value VARCHAR(50) NOT NULL DEFAULT '',
    PRIMARY KEY (api_setting_id, api_param_setting_no),
    FOREIGN KEY (api_setting_id) REFERENCES ApiSetting (api_setting_id)
);

CREATE TABLE IF NOT EXISTS ApiResultHistory (
    api_result_id INT NOT NULL AUTO_INCREMENT,
    api_id INT NOT NULL,
    request_endpoint VARCHAR(50) NOT NULL DEFAULT '',
    request_param_string VARCHAR(1000) NOT NULL DEFAULT '',
    request_date_time DATETIME NOT NULL,
    response_status_code VARCHAR(3) NOT NULL,
    response_data TEXT NOT NULL DEFAULT (''),
    response_date_time DATETIME NOT NULL,
    PRIMARY KEY (api_result_id),
    FOREIGN KEY (api_id) REFERENCES Api (api_id)
);
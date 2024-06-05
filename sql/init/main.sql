CREATE TABLE devices (
	id INT AUTO_INCREMENT PRIMARY KEY,
	mac_address TEXT NOT NULL,
	device_type TEXT NOT NULL,
	control_flg TINYINT(1) NOT NULL DEFAULT 1,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE devices_controller (
	id INT AUTO_INCREMENT PRIMARY KEY,
	device_id INT NOT NULL,
	device_type INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	Foreign Key (device_id) REFERENCES devices(id)
);

CREATE TABLE control_records (
	id INT AUTO_INCREMENT PRIMARY KEY,
	control_record TEXT NOT NULL,
	device_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	Foreign Key (device_id) REFERENCES devices(id)
);

CREATE TABLE error_message_records (
	id INT AUTO_INCREMENT PRIMARY KEY,
	device_id INT,
	error_type INT NOT NULL,
	message TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	Foreign Key (device_id) REFERENCES devices(id)
);

CREATE TABLE device_address_records (
	id INT AUTO_INCREMENT PRIMARY KEY,
	address CHAR(16) NOT NULL,
	device_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	Foreign Key (device_id) REFERENCES devices(id)
);

CREATE TABLE status_records (
	id INT AUTO_INCREMENT PRIMARY KEY,
	device_id INT,
	status INT(3) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	Foreign Key (device_id) REFERENCES devices(id)
);

package com.pinkpanther.java_app;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.mongodb.repository.config.EnableMongoRepositories;

import com.pinkpanther.java_app.repository.EmployeeRepository;

@SpringBootApplication
@EnableMongoRepositories
public class JavaAppApplication {

	// @Autowired
    // EmployeeRepository employeeRepo;

	public static void main(String[] args) {
		SpringApplication.run(JavaAppApplication.class, args);
	}

}

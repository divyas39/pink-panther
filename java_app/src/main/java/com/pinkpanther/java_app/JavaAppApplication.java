package com.pinkpanther.java_app;

// import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.mongodb.repository.config.EnableMongoRepositories;
// import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@SpringBootApplication
@EnableMongoRepositories
@EnableJpaRepositories
public class JavaAppApplication {

	// @Autowired`
    // EmployeeRepository employeeRepo;

	public static void main(String[] args) {
		SpringApplication.run(JavaAppApplication.class, args);
	}

}

package com.pinkpanther.java_app.repository;

import java.util.Optional;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

import com.pinkpanther.java_app.model.Employee;

public interface EmployeeRepository extends MongoRepository<Employee, String> {
    
    @Query("{id:'?0'}")
    Optional<Employee> findEmployeeById(String id);

    public long count();

}

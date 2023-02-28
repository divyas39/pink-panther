package com.pinkpanther.java_app.repository;

import java.util.List;
import java.util.Optional;
// import java.util.Calendar;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;

import com.pinkpanther.java_app.model.Employee;

public interface EmployeeRepository extends MongoRepository<Employee, String> {
    
    // @Query("{name:'?0', experience:'?0'}")
    // Employee createItem(String id, String name, int experience);
    
    @Query("{id:'?0'}")
    Optional<Employee> findEmployeeById(String id);
    
    // @Query(fields="{'name' : 1, 'experience' : 1, 'dob' : 1}")
    // @Query("")
    // List<Employee> findAll();

    @Query("{id:'?0', name:'?0', experience:'?0', dob:'?0'}")
    Employee updateEmployee(String id, String name, int experience);

    // @Query("{id:'?0'}")
    // Employee deleteEmployeeById(String id);

    // @Query(fields="{'id' : 1}")
    // void deleteAll();
    
    public long count();

}

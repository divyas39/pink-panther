package com.pinkpanther.java_app.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.pinkpanther.java_app.model.Employee;


public interface SqldbRepository extends JpaRepository<Employee, String>{
    
}




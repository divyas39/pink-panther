package com.pinkpanther.java_app.controller;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

import com.pinkpanther.java_app.model.Employee;
import com.pinkpanther.java_app.repository.EmployeeRepository;

@CrossOrigin(origins = "http://localhost:8080")
@RestController
@RequestMapping("/api")
public class Controller{

    @Autowired
    EmployeeRepository employeeRepo;

    @PostMapping("/employees")
    public ResponseEntity<?> createEmployees(@RequestBody Employee employee){
        try{
            Employee _employee = employeeRepo.save(new Employee(employee.getName(), employee.getExperience()));
            return new ResponseEntity<>(_employee, HttpStatus.CREATED);
        }

        catch (Exception e){
            return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/employees")
    public ResponseEntity<List<Employee>> getAllEmployees() {
        try {
            List<Employee> employees = new ArrayList<Employee>();
            employees = employeeRepo.findAll();

            if (employees.isEmpty()) {
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
            }

            return new ResponseEntity<>(employees, HttpStatus.OK);
        } 
            catch (Exception e) {
                return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @GetMapping("/employee/{id}")
    public ResponseEntity<Employee> getEmployeeById(@PathVariable("id") String id) {
       Optional<Employee> employeeData = employeeRepo.findEmployeeById(id);

        if (employeeData.isPresent()) {
        return new ResponseEntity<>(employeeData.get(), HttpStatus.OK);
        } 
        
        else {
        return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
    }
    

    @PutMapping("/employee/{id}")
    public ResponseEntity<Employee> updateEmployee(@PathVariable("id") String id, @RequestBody Employee employee){
        Optional<Employee> oldEmployee = employeeRepo.findEmployeeById(id);
        if(oldEmployee.isPresent()) {
            Employee newEmployee = oldEmployee.get();
            newEmployee.setName(employee.getName());
            newEmployee.setExperience(employee.getExperience());
            return new ResponseEntity<>(employeeRepo.save(newEmployee), HttpStatus.OK);
        }

        else {
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        }
    }

    @DeleteMapping("/employee/{id}")
    public ResponseEntity<HttpStatus> deleteEmployee(@PathVariable("id") String id) {
        try {
        employeeRepo.deleteById(id);
        return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
        return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }

    @DeleteMapping("/employees")
    public ResponseEntity<HttpStatus> deleteAllEmployees() {
        try {
        employeeRepo.deleteAll();
        return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        } catch (Exception e) {
        return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
}
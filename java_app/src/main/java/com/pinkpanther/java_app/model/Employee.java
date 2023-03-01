package com.pinkpanther.java_app.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Table;
// import jakarta.persistence.Id;
// import java.util.Calendar;

@Entity
@Table(name = "employees")
@Document("employees")
public class Employee {
    @Id
    private String id;

    private String name;
    private int experience;
    // private Calendar dob;

    public Employee(){

    }

    public Employee(String name, int experience){
        // super();
        // this.id = id;
        this.name = name;
        this.experience = experience;
        // this.dob = dob;
    }

    // @Id
    // @GeneratedValue(strategy = GenerationType.IDENTITY)
    public String getId() {
        return id;  
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getExperience() {
        return experience;
    }

    public void setExperience(int experience) {
        this.experience = experience;
    }
}

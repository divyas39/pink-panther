package com.pinkpanther.java_app.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
// import java.util.Calendar;

@Document("employees")
public class Employee {
    @Id
    private String id;

    private String name;
    private int experience;
    // private Calendar dob;

    public String getId() {
        return id;  
    }

    // public void setId(String id) {
    //     this.id = id;
    // }

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

    // public Calendar getDob() {
    //     return dob;
    // }

    // public void setDob(Calendar dob) {
    //     this.dob = dob;
    // }


    public Employee(String name, int experience){
        super();
        // this.id = id;
        this.name = name;
        this.experience = experience;
        // this.dob = dob;
   }
}

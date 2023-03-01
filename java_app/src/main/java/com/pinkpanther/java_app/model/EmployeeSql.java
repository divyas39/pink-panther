// package com.pinkpanther.java_app.model;

// import jakarta.persistence.Entity;
// import jakarta.persistence.GeneratedValue;
// import jakarta.persistence.GenerationType;
// import jakarta.persistence.Table;
// import jakarta.persistence.Id;
// // import jakarta.persistence.Column;

// // import java.util.Calendar;
// // import javax.persistence.Entity;
// // import javax.persistence.GeneratedValue;
// // import javax.persistence.GenerationType;
// // import javax.persistence.Id;


// @Entity
// @Table(name = "employees")
// public class EmployeeSql {
//     // @Id 
//     // @GeneratedValue(strategy = GenerationType.IDENTITY)
//     private String id;

//     // @Column(nullable = false)
//     private String name;

//     // @Column(nullable = false)
//     private int experience;
//     // private Calendar dob;

//     public EmployeeSql(){

//     }

//     public EmployeeSql(String name, int experience){
//         // super();
//         // this.id = id;
//         this.name = name;
//         this.experience = experience;
//         // this.dob = dob;
//    }

//     @Id
//     @GeneratedValue(strategy = GenerationType.IDENTITY)
//     public String getId() {
//         return id;  
//     }

//     public void setId(String id) {
//         this.id = id;
//     }

//     public String getName() {
//         return name;
//     }

//     public void setName(String name) {
//         this.name = name;
//     }

//     public int getExperience() {
//         return experience;
//     }

//     public void setExperience(int experience) {
//         this.experience = experience;
//     }

//     // public Calendar getDob() {
//     //     return dob;
//     // }

//     // public void setDob(Calendar dob) {
//     //     this.dob = dob;
//     // }
// }

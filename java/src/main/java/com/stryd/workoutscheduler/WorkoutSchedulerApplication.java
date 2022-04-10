package com.stryd.workoutscheduler;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class WorkoutSchedulerApplication {

	public static void main(String[] args) {
		SpringApplication.run(WorkoutSchedulerApplication.class, args);
	}


	@GetMapping("/")
	public String Hello() {
		return "<p>Hello, World!</p>";
	}

	// Add your new REST endpoint(s) here
}

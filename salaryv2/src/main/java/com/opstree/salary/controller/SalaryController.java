package com.opstree.salary.controller;

import com.opstree.salary.model.SalaryModel;
import com.opstree.salary.service.SalaryService;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController()
public class SalaryController {

    private SalaryService service;

    @Autowired
    public SalaryController(SalaryService service) {

        this.service = service;
    }

    @GetMapping("/health")
    public String test(){

        return "Success";
    }

    @PostMapping("/profiles")
    public ResponseEntity createProfile(@RequestBody SalaryModel document) throws Exception {

        return new ResponseEntity(service.createSalaryModel(document), HttpStatus.CREATED);
    }

    @PutMapping
    public ResponseEntity updateProfile(@RequestBody SalaryModel document) throws Exception {

        return new ResponseEntity(service.updateProfile(document), HttpStatus.CREATED);
    }

    @GetMapping("/{id}")
    public SalaryModel findById(@PathVariable String id) throws Exception {

        return service.findById(id);
    }

    @GetMapping
    public List<SalaryModel> findAll() throws Exception {

        return service.findAll();
    }

    @GetMapping(value = "/search")
    public List<SalaryModel> search(@RequestParam(value = "technology") String technology) throws Exception {
        return service.searchByTechnology(technology);
    }

    @GetMapping(value = "/api/v1/profiles/name-search")
    public List<SalaryModel> searchByName(@RequestParam(value = "name") String name) throws Exception {
        return service.findProfileByName(name);
    }


    @DeleteMapping("/{id}")
    public String deleteSalaryModel(@PathVariable String id) throws Exception {

        return service.deleteSalaryModel(id);

    }
}

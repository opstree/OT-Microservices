package com.opstree.salary;

import static org.hamcrest.Matchers.containsString;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.print;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import org.junit.Before;
import org.junit.Test;
import static org.junit.Assert.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.test.web.servlet.RequestBuilder;
import org.springframework.web.context.WebApplicationContext;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

public class TestController extends ApplicationTest {

	@Autowired
	private MockMvc mockMvc;

	@Test
	public void testHealthCheck() throws Exception {
		RequestBuilder request = MockMvcRequestBuilders
				.get("/health");

		MvcResult result = mockMvc.perform(request)
				.andExpect(status().isOk())
				.andExpect(content().string("Success"))
				.andReturn();
		assertEquals("Success", result.getResponse().getContentAsString());
	}
}

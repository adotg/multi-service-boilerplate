package boilerplate.multiservice.data;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.request.MockHttpServletRequestBuilder;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import org.springframework.web.context.WebApplicationContext;

public abstract class AbstractTest {
    @Autowired
    WebApplicationContext webApplicationContext;

    protected MockMvc mvc;

    protected void setUp() {
        mvc = MockMvcBuilders.webAppContextSetup(webApplicationContext).build();
    }

    protected String sendRequest(String uri, HttpMethod method, String body) throws Exception {
        MockHttpServletRequestBuilder requestBuilder;
        if (method == HttpMethod.POST) {
            requestBuilder = MockMvcRequestBuilders.post(uri)
                .accept(MediaType.APPLICATION_JSON)
                .contentType(MediaType.APPLICATION_JSON)
                .content(body);
        } else {
            requestBuilder = MockMvcRequestBuilders.get(uri)
                .accept(MediaType.APPLICATION_JSON);
        }

        MvcResult mvcResult = mvc.perform(requestBuilder).andReturn();
        return mvcResult.getResponse().getContentAsString();
    }
}

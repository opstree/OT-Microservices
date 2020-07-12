import react, * as React from "react";
import { Page, Grid } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";
import { Button, Form, FormGroup, Label, Input } from 'reactstrap';
import { withFormik } from 'formik';

const EmployeeForm = ({ values, handleChange, handleSubmit, errors, touched, isSubmitting }) => {
  return (
    <SiteWrapper>
      <Page.Card
            title="Employee Registration"
        ></Page.Card>
        <Grid.Col md={6} lg={6} className="align-self-center">
        <Form onSubmit={handleSubmit}>
          <FormGroup>
            {touched.id && errors.id && <p className="red">{errors.id}</p>}
            <Label for="id">Employee ID</Label>
            <Input 
              type="text" 
              name="id"
              value={values.id}
              onChange={handleChange}
              id="id" 
              placeholder="Employee ID" 
            />
          </FormGroup>
          <FormGroup>
            {touched.name && errors.name && <p className="red">{errors.name}</p>}
            <Label for="name">Name</Label>
            <Input 
              type="text" 
              name="name"
              value={values.name}
              onChange={handleChange}
              id="name" 
              placeholder="Employee Name" 
            />
          </FormGroup>
          <FormGroup>
            {touched.address && errors.address && <p className="red">{errors.address}</p>}
            <Label for="address">Address</Label>
            <Input 
              type="text" 
              name="address"
              value={values.address}
              onChange={handleChange}
              id="address" 
              placeholder="Employee Address" 
            />
          </FormGroup>
          <FormGroup>
            {touched.email && errors.email && <p className="red">{errors.email}</p>}
            <Label for="email">Email ID</Label>
            <Input 
              type="text" 
              name="email"
              value={values.email}
              onChange={handleChange}
              id="email" 
              placeholder="Email ID" 
            />
          </FormGroup>
          <FormGroup>
            {touched.phone_number && errors.phone_number && <p className="red">{errors.phone_number}</p>}
            <Label for="phone_number">Phone Number</Label>
            <Input 
              type="text" 
              name="phone_number"
              value={values.phone_number}
              onChange={handleChange}
              id="phone_number" 
              placeholder="Phone Number" 
            />
          </FormGroup>
          <FormGroup>
            {touched.annual_package && errors.annual_package && <p className="red">{errors.annual_package}</p>}
            <Label for="annual_package">Annual Package</Label>
            <Input 
              type="number" 
              name="annual_package"
              value={values.annual_package}
              onChange={handleChange}
              id="annual_package" 
              placeholder="Annual Package" 
            />
          </FormGroup>
          <FormGroup>
            {touched.job_role && errors.job_role && <p className="red">{errors.job_role}</p>}
            <Label for="job_role">Job Role</Label>
            <Input type="select" name="job_role" id="job_role" value={values.job_role} onChange={handleChange}>
              <option>Select Role</option>
              <option>Developer</option>
              <option>DevOps</option>
            </Input>
          </FormGroup>
          <FormGroup>
            {touched.status && errors.status && <p className="red">{errors.status}</p>}
            <Label for="status">Status</Label>
            <Input type="select" name="status" id="status" value={values.status} onChange={handleChange}>
              <option>Select Status</option>
              <option>Ex-Employee</option>
              <option>Current Employee</option>
            </Input>
          </FormGroup>
          <FormGroup>
            {touched.location && errors.location && <p className="red">{errors.location}</p>}
            <Label for="location">Location</Label>
            <Input type="select" name="location" id="location" value={values.location} onChange={handleChange}>
              <option>Select Location</option>
              <option>Delhi</option>
              <option>Bangalore</option>
              <option>Hyderabad</option>
              <option>Newyork</option>
            </Input>
          </FormGroup>
          <FormGroup>
            {touched.joining_date && errors.joining_date && <p className="red">{errors.joining_date}</p>}
            <Label for="joining_date">Joining Date</Label>
            <Input
              type="date"
              name="joining_date"
              id="joining_date"
              placeholder="datetime placeholder"
              value={values.joining_date} 
              onChange={handleChange}
            />
          </FormGroup>
          <Button color="primary" disabled={isSubmitting}>Submit</Button>
        </Form>
    </Grid.Col>
    </SiteWrapper>
  );
}

const FormikApp = withFormik({
  mapPropsToValues({ username, password }) {
    return { username, password }
  },
  handleSubmit(values, { props, resetForm, setErrors, setSubmitting }) {
    fetch('/employee/create', {
      method: 'POST',
      body: JSON.stringify(values),
      headers: {
          'Content-Type': 'application/json'
    }})
    fetch('/notification/send', {
      method: 'POST',
      body: JSON.stringify(values),
      headers: {
          'Content-Type': 'application/json'
    }})
    console.log(JSON.stringify(values))
  }
})(EmployeeForm);

export default FormikApp

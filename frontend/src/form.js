import React  from 'react';
import { Formik } from 'formik'
import { Form, Input, SubmitBtn, Select, Datepicker } from 'react-formik-ui';
import CssBaseline from "@material-ui/core/CssBaseline";
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import { makeStyles } from '@material-ui/core/styles';

const gatewayURL = process.env.REACT_APP_GATEWAY_URL

const employeeURL = gatewayURL + "/create"

const useStyles = makeStyles((theme) => ({ 
  fixedHeight: {
      height: 350,
    },
    appBarSpacer: theme.mixins.toolbar,
    content: {
      flexGrow: 1,
      height: '100vh',
      overflow: 'auto',
    },
    container: {
      paddingTop: theme.spacing(8),
      paddingBottom: theme.spacing(4),
    },
}));


export default function UserForm() {
  const classes = useStyles();
  return (
    <div>
      <CssBaseline />
      <Container maxWidth="lg" className={classes.container}>
      <Grid>
        <Formik
          initialValues={{
            id: '',
            name: '',
            job_role: '',
            joining_date: '',
            address: '',
            city: '',
            email_id: '',
            annual_package: '',
            phone_number: ''
          }}
          onSubmit={data =>
            fetch(employeeURL, {
              method: 'POST',
              body: JSON.stringify(data),
              headers: {
                  'Content-Type': 'application/json'
            }})
            (alert(JSON.stringify(data)))
          }
        >
          <Form mode='themed'>

            <Input
              name='id'
              label='Employee ID'
              placeholder='1'
              required
            />

            <Input
              name='name'
              label='Name'
              placeholder='Abhishek'
              required
            />
            <Select
              name='city'
              label='State'
              placeholder='Select State'
              options={[
                { value: 'Delhi', label: 'Delhi' },
                { value: 'Himachal Pradesh', label: 'Himachal Pradesh' },
                { value: 'Karnataka', label: 'Karnataka' },
              ]}
            />
            <Select
              name='job_role'
              label='Role'
              placeholder='Select Role'
              options={[
                { value: 'DevOps', label: 'DevOps' },
                { value: 'Developer', label: 'Developer' },
              ]}
            />
            <Input
              name='address'
              label='Address'
              placeholder='A-27 D, Noida'
              required
            />
            <Datepicker
              name='joining_date'
              label='Joining Date'
              placeholder='DD.MM.YYYYY'
              dateFormat='dd.MM.yyyy'
            />
            <Input
              name='email_id'
              label='Email ID'
              placeholder='abhishek@example.com'
              required
            />
            <Input
              name='annual_package'
              label='Annual Package'
              placeholder='10000'
              type='number'
              required
            />
            <Input
              name='phone_number'
              label='Phone Number'
              required
            />
            <SubmitBtn />
          </Form>
        </Formik>
        </Grid>
      </Container>
    </div>
  )
}

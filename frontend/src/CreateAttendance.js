import React  from 'react';
import { Formik } from 'formik'
import { Form, Input, SubmitBtn, Select, Datepicker } from 'react-formik-ui';
import CssBaseline from "@material-ui/core/CssBaseline";
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import { makeStyles } from '@material-ui/core/styles';

const gatewayURL = process.env.REACT_APP_GATEWAY_URL

const attendanceURL = gatewayURL + "/attendance/create"

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
            fetch(attendanceURL, {
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
              type='number'
              required
            />
            <Select
              name='status'
              label='Status'
              placeholder='Select status'
              options={[
                { value: 'Present', label: 'Present' },
                { value: 'Absent', label: 'Absent' },
              ]}
            />
            <Datepicker
              name='date'
              label='Date'
              placeholder='DD.MM.YYYYY'
              dateFormat='dd.MM.yyyy'
            />
            <SubmitBtn />
          </Form>
        </Formik>
        </Grid>
      </Container>
    </div>
  )
}

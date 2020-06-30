import React from 'react';
import { Formik, Field, Form, ErrorMessage } from 'formik';
import { withStyles } from '@material-ui/core/styles';
import CssBaseline from "@material-ui/core/CssBaseline";
import Container from '@material-ui/core/Container';
import * as Yup from 'yup';

const styles = theme => ({ 
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
});


class UserForm extends React.Component {
    render() {
      const { classes } = this.props;
        return (
          <div>
          <CssBaseline />
          <Container maxWidth="lg" className={classes.container}>
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
                validationSchema={Yup.object().shape({
                    id: Yup.string()
                        .required('Employee ID is required'),
                    name: Yup.string()
                        .required('Name is required'),
                    email_id: Yup.string()
                        .email('Email ID is invalid')
                        .required('Email ID is required'),
                    job_role: Yup.string()
                        .required('Job Role is required'),
                    joining_date: Yup.string()
                        .required('Joining date is required'),
                    address: Yup.string()
                        .required('Address is required'),
                    city: Yup.string()
                        .required('City is required'),
                    annual_package: Yup.number()
                      .required('Annual Package is required'),
                    phone_number: Yup.string()
                      .required('Phone number is required')
                })}
                onSubmit={fields => {
                    alert('SUCCESS!! :-)\n\n' + JSON.stringify(fields, null, 4))
                }}
                render={({ errors, status, touched }) => (
                    <Form>
                        <div className="form-group">
                            <label htmlFor="id">Employee ID</label>
                            <Field name="id" type="text" className={'form-control' + (errors.id && touched.id ? ' is-invalid' : '')} />
                            <ErrorMessage name="id" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="name">Name</label>
                            <Field name="name" type="text" className={'form-control' + (errors.name && touched.name ? ' is-invalid' : '')} />
                            <ErrorMessage name="name" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="email_id">Email</label>
                            <Field name="email_id" type="text" className={'form-control' + (errors.email_id && touched.email_id ? ' is-invalid' : '')} />
                            <ErrorMessage name="email_id" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="job_role">Job Role</label>
                            <Field name="job_role" type="text" className={'form-control' + (errors.job_role && touched.job_role ? ' is-invalid' : '')} />
                            <ErrorMessage name="job_role" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="address">Address</label>
                            <Field name="address" type="text" className={'form-control' + (errors.address && touched.address ? ' is-invalid' : '')} />
                            <ErrorMessage name="address" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="joining_date">Joining Date</label>
                            <Field name="joining_date" type="text" className={'form-control' + (errors.joining_date && touched.joining_date ? ' is-invalid' : '')} />
                            <ErrorMessage name="joining_date" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="city">City</label>
                            <Field name="city" type="text" className={'form-control' + (errors.city && touched.city ? ' is-invalid' : '')} />
                            <ErrorMessage name="city" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="phone_number">Phone Number</label>
                            <Field name="phone_number" type="text" className={'form-control' + (errors.phone_number && touched.phone_number ? ' is-invalid' : '')} />
                            <ErrorMessage name="phone_number" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <label htmlFor="annual_package">Annual Package</label>
                            <Field name="annual_package" type="number" className={'form-control' + (errors.annual_package && touched.annual_package ? ' is-invalid' : '')} />
                            <ErrorMessage name="annual_package" component="div" className="invalid-feedback" />
                        </div>
                        <div className="form-group">
                            <button type="submit" className="btn btn-primary mr-2">Register</button>
                            <button type="reset" className="btn btn-secondary">Reset</button>
                        </div>
                    </Form>
                )}
            />
            </Container>
            </div>
        )
    }
}

// export default withStyles(styles, { withTheme: true })(SignupForm);

export default withStyles(styles, { withTheme: true })(UserForm);

// import React from 'react';
// import { useFormik } from 'formik';
// import { fade, withStyles } from '@material-ui/core/styles';
// import TextField from '@material-ui/core/TextField';
// import Container from '@material-ui/core/Container';
// import Button from '@material-ui/core/Button';

// import Input from '@material-ui/core/Input';

// const styles = theme => ({ 
//   fixedHeight: {
//       height: 350,
//     },
//     appBarSpacer: theme.mixins.toolbar,
//     content: {
//       flexGrow: 1,
//       height: '100vh',
//       overflow: 'auto',
//     },
//     container: {
//       paddingTop: theme.spacing(8),
//       paddingBottom: theme.spacing(4),
//     },
// });


// class SignupForm extends React.Component {

//   constructor(props) {
//     super(props);
//     this.state = {value: ''};

//     this.handleChange = this.handleChange.bind(this);
//     this.handleSubmit = this.handleSubmit.bind(this);
//   }

//   handleSubmit(event) {
//     event.preventDefault();
//     const data = new FormData(event.target);
//     console.log(JSON.stringify(data))
//     // alert(JSON.stringify(event, null, 2)); 
//     alert(JSON.stringify(data));
//   }

//   handleChange(event) {
//     this.setState({value: event.target.value});
//   }
//   render() {
//     const { classes } = this.props;
//     return (
//       <Container maxWidth="lg" className={classes.container}>
//         <form onSubmit={this.handleSubmit}>
//         <div>
//          <TextField label="Name" type="text" name="name" onChange={this.handleChange} />
//          </div>
//          <div>
//          <Button variant="contained">
//             <Input type="submit">
//                 Submit
//             </Input>
//           </Button>
//          </div>
//           {/* <button type="submit">Submit</button> */}
//         </form>
//         </Container>
//       );
//   }
// };

// export default withStyles(styles, { withTheme: true })(SignupForm);

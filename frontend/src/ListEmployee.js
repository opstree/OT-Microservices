import React, { useState, useEffect } from 'react';
import MaterialTable from 'material-table';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from "@material-ui/core/CssBaseline";
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import clsx from 'clsx';

const useStyles = makeStyles((theme) => ({ 
    fixedHeight: {
        height: 600,
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
      paper: {
        padding: theme.spacing(2),
        display: 'flex',
        overflow: 'auto',
        flexDirection: 'column',
      },
}));

const columns = [
    { title: 'ID', field: 'id' },
    { title: 'Name', field: 'name' },
    { title: 'Job Role', field: 'job_role' },
    { title: 'Joining Date', field: 'joining_date' },
]

export default function EmployeeList() {
    const [stats, handleStats] = useState([]);

    const FetchData = async () => {
      const data = await fetch('http://172.17.0.3:8080/search/all');
      const stats = await data.json();
      handleStats(stats)  
    }
  
    useEffect(() => {
      FetchData()
    }, [])

    const classes = useStyles();
    const empData = stats
    return (
        <div>
        <CssBaseline />
        <Container className={classes.container}>
        {/* <Grid container spacing={3}> */}
            <MaterialTable
            title="Employees"
            columns={columns}
            key={empData.id}
            data={empData}
            />
        {/* </Grid> */}
        </Container>
    </div>
    )
}

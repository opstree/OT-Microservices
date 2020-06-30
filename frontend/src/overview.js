import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CssBaseline from "@material-ui/core/CssBaseline";
import Container from '@material-ui/core/Container';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import clsx from 'clsx';
import Deposits from './employeeCount.js';
import Charts from './jobRole.js';
import Map from './map.js';

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
      paper: {
        padding: theme.spacing(2),
        display: 'flex',
        overflow: 'auto',
        flexDirection: 'column',
      },
}));

export default function Overview() {
    const classes = useStyles();
    const fixedHeightPaper = clsx(classes.paper, classes.fixedHeight);
    return (
        <div>
            <CssBaseline />
            <Container maxWidth="lg" className={classes.container}>
            <Grid container spacing={3}>
                {/* Chart */}
                <Grid item xs={12} md={10} lg={9} >
                <Paper className={fixedHeightPaper}>
                    <Charts />
                </Paper>
                </Grid>
                {/* Recent Deposits */}
                <Grid item xs={12} md={4} lg={3} >
                {/* } */}
                <Paper className={fixedHeightPaper}>
                    <Deposits />
                </Paper>
                </Grid>
                {/* Recent Orders */}
                <Map /> 
            </Grid>
            </Container>
        </div>
    );
}

import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Title from './Title';

const useStyles = makeStyles({
  depositContext: {
    flex: 1,
    paddingTop: 30,
  },
});

export default function Deposits() {
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
  const date =  new Date().toLocaleString();
  const data2 = stats.length
  return (
    <React.Fragment>
      <Typography align="center">
      <Title>Total Employees</Title>
      </Typography>
      <Typography component="p" variant="h1" align="center" className={classes.depositContext}>
      {data2}
      </Typography>
      <Typography color="textSecondary" className={classes.depositContext} align="center">
      {date}
      </Typography>
    </React.Fragment>
  );
}


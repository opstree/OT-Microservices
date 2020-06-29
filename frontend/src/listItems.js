import React from 'react';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListSubheader from '@material-ui/core/ListSubheader';
import { makeStyles } from '@material-ui/core/styles';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import Collapse from '@material-ui/core/Collapse';
import GroupIcon from '@material-ui/icons/Group';
import HowToRegIcon from '@material-ui/icons/HowToReg';
import MonetizationOnIcon from '@material-ui/icons/MonetizationOn';
import CreateIcon from '@material-ui/icons/Create';
import ListIcon from '@material-ui/icons/List';
import { Router, Link } from 'react-router-dom';
import { createBrowserHistory } from "history";
import SignupForm from './form.js';

const history = createBrowserHistory();

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    maxWidth: 360,
    backgroundColor: theme.palette.background.paper,
  },
  nested: {
    paddingLeft: theme.spacing(4),
  },
}));

export default function MainListItems() {
  const classes = useStyles();
  const [open, setOpen] = React.useState(true);
  const handleClick = () => {
    setOpen(!open);
  };

  const [attendopen, attendsetOpen] = React.useState(true);
  const attendhandleClick = () => {
    attendsetOpen(!attendopen);
  };

  return (
    <Router history={history}>
    <List
      component="nav"
      aria-labelledby="nested-list-subheader"
      subheader={
        <ListSubheader component="div" id="nested-list-subheader">
          Options
        </ListSubheader>
      }
      className={classes.root}
    >
      <ListItem component={Link} to="/">
        <ListItemIcon>
          <MonetizationOnIcon />
        </ListItemIcon>
        <ListItemText primary="Overview"/>
      </ListItem>
      <ListItem button onClick={handleClick}>
        <ListItemIcon>
          <GroupIcon />
        </ListItemIcon>
        <ListItemText primary="Employees" />
        {open ? <ExpandLess /> : <ExpandMore />}
      </ListItem>
      <Collapse in={open} timeout="auto" unmountOnExit>
        <List component="div" disablePadding>
          <ListItem button className={classes.nested} component={Link} to="/#/create">
            <ListItemIcon>
              <CreateIcon />
            </ListItemIcon>
            <ListItemText primary="Add Data" />
          </ListItem>
          <ListItem button className={classes.nested}>
            <ListItemIcon>
              <ListIcon />
            </ListItemIcon>
            <ListItemText primary="List Data" />
          </ListItem>
        </List>
      </Collapse>
      <ListItem button onClick={attendhandleClick}>
        <ListItemIcon>
          <HowToRegIcon />
        </ListItemIcon>
        <ListItemText primary="Attendance" />
        {attendopen ? <ExpandLess /> : <ExpandMore />}
      </ListItem>
      <Collapse in={attendopen} timeout="auto" unmountOnExit>
        <List component="div" disablePadding>
          <ListItem button className={classes.nested}>
            <ListItemIcon>
              <CreateIcon />
            </ListItemIcon>
            <ListItemText primary="Add Data" />
          </ListItem>
        </List>
      </Collapse>
      <ListItem>
        <ListItemIcon>
          <MonetizationOnIcon />
        </ListItemIcon>
        <ListItemText primary="Salary" />
      </ListItem>
    </List>
    </Router>
  );
}

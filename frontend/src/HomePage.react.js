// @flow

import * as React from "react";

import {
  Page,
  Avatar,
  Icon,
  Grid,
  Card,
  Text,
  Table,
  Alert,
  Progress,
  colors,
  Dropdown,
  Button,
  StampCard,
  StatsCard,
  ProgressCard,
  Badge,
} from "tabler-react";

import C3Chart from "react-c3js";

import SiteWrapper from "./SiteWrapper.react";
import { ListAllEmployees, 
  ListEmployeeActiveEmployee,
  ListEmployeeInActiveEmployee,
  RoleDistribution,
  LocationDistribution,
  StatusDistribution } from './EmployeeData';

function Home() {
  return (
    <SiteWrapper>
      <Page.Content title="Dashboard">
        <Grid.Row cards={true}>
          <ListAllEmployees/>
          <ListEmployeeActiveEmployee/>
          <ListEmployeeInActiveEmployee/>
          <Grid.Col sm={3}>
            <StatsCard
              layout={1}
              movement={0}
              total="4"
              label="Office Locations"
            />
          </Grid.Col>
          <Grid.Col>
            <Grid.Row cards="true">
              <RoleDistribution/>
              <StatusDistribution/>
              <LocationDistribution/>
            </Grid.Row>
          </Grid.Col>
        </Grid.Row>
      </Page.Content>
    </SiteWrapper>
  );
}

export default Home;
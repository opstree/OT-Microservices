import * as React from "react";
import { useState, useEffect } from 'react';
import { DonutChart } from "@opd/g2plot-react";

require('dotenv').config();

const gatewayURL = process.env.REACT_APP_GATEWAY_URL;

const employeeURL = gatewayURL + "/employee/search/roles"

export default function Chart() {
  const [stats, handleStats] = useState([]);

  const FetchData = async () => {
    const data = await fetch(employeeURL);
    const stats = await data.json();
    handleStats(stats)  
  }
  const data = stats
  console.log(stats)

  const config = {
    forceFit: true,
    title: {
      visible: true,
      text: "Roles Distribution",
      style:{
        fontSize: 15,
        fill: 'black',
      }
    },
    statistic: {
      totalLabel: "Total"
    },
    description: {
      visible: false,
      text: "环图指标卡能够代替tooltip，在环图中心挖空部分显示各分类的详细信息。"
    },
    radius: 0.9,
    height: 300,
    padding: "auto",
    data,
    angleField: "value",
    colorField: "type"
  };

  useEffect(() => {
    FetchData()
  }, [])
  return (
    <section>
    <DonutChart {...config} />
  </section>
  )
}

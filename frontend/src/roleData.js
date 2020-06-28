import { useState, useEffect } from 'react';

function GetJobRoles() {
    const [stats, handleStats] = useState([]);
  
    const FetchData = async () => {
      const data = await fetch('http://172.17.0.3:8080/search/all');
      const stats = await data.json();
      handleStats(stats)  
    }
  
    useEffect(() => {
      FetchData()
    }, [])
    return stats
  }
  
export default DataGeneration() {
    const stats = GetJobRoles()
    var arr = stats, obj = {};
    for (var i = 0; i < arr.length; i++) {
      if (!obj[arr[i].job_role]) {
        obj[arr[i].job_role] = 1;
      } else if (obj[arr[i].job_role]) {
        obj[arr[i].job_role] += 1;
      }
    }
    const data = {
      labels: ["DevOps", "Developer", "Manager", "HR"],
      datasets: [
        {
          data: [obj["DevOps"], obj["Develeper"], obj["Manager"], obj["HR"]],
          backgroundColor: ["#949FB1", "#46BFBD", "#FDB45C", "#4D5360", "#4D5360"],
          hoverBackgroundColor: [
            "#FF5A5E",
            "#5AD3D1",
            "#FFC870",
            "#A8B3C5",
            "#616774"
          ]
        }
      ]
    }
    return data
}

import React from "react";
import { Doughnut } from "react-chartjs-2";
import { MDBContainer } from "mdbreact";

function preventDefault(event) {
    event.preventDefault();
}
  
class Charts extends React.Component {
state = {
  dataDoughnut: {
    labels: ["DevOps", "Human Resources", "Manager"],
    datasets: [
      {
        data: [300, 50, 100],
        backgroundColor: ["#949FB1", "#46BFBD", "#FDB45C", "#949FB1", "#4D5360"],
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
}

render() {
    return (
    <MDBContainer>
      <h6 className="mt-5">Job Role Distribution</h6>
      <Doughnut data={this.state.dataDoughnut} options={{ responsive: true, maintainAspectRatio: true }} height="50" width="180"/>
    </MDBContainer>
    );
  }
}

export default Charts;

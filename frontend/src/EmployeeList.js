import react, * as React from "react";
import { Page, Grid, Table, Button } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";
//import { withTransaction } from '@elastic/apm-rum-react';

class ListEmployee extends React.Component {
	constructor(props) {
		super(props)
		this.state = { data: [] }
	}
	
	loadData() {
		fetch("/employee/search/all")
			.then(response => response.json())
			.then(data => {
				this.setState({data: data })
		})
			.catch(err => console.error(this.props.url, err.toString()))
	}

	componentDidMount() {
		this.loadData()
	}
	
  render() {
      return (
          <SiteWrapper>
          <Page.Card
              title="Employee List"
          ></Page.Card>
          <Grid.Col md={6} lg={10} className="align-self-center">
          <Table>
            <Table.Header>
                 <Table.ColHeader>Employee ID</Table.ColHeader>
                 <Table.ColHeader>Name</Table.ColHeader>
                 <Table.ColHeader>Email</Table.ColHeader>
                 <Table.ColHeader>Phone Number</Table.ColHeader>
                 <Table.ColHeader>Job Role</Table.ColHeader>
                 <Table.ColHeader>Job Location</Table.ColHeader>
            </Table.Header>
            <Table.Body>
           { this.state.data.map((item, i) => {
                return (
                    <Table.Row>
                        <Table.Col>{item.id}</Table.Col>
                        <Table.Col>{item.name}</Table.Col>
                        <Table.Col>{item.email}</Table.Col>
                        <Table.Col>{item.phone_number}</Table.Col>
                        <Table.Col>{item.job_role}</Table.Col>
                        <Table.Col>{item.location}</Table.Col>
                    </Table.Row>  
                );
                })  
            }
            </Table.Body>
            </Table>
          </Grid.Col>
          </SiteWrapper>
      );
  }
}

export default ListEmployee
//export default withTransaction('ListEmployee', 'component')(ListEmployee)
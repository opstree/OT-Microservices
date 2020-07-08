import react, * as React from "react";
import { Page, Grid, Table, Button } from "tabler-react";
import SiteWrapper from "./SiteWrapper.react";
import { Component } from 'react';
import Doc from './DocService';
import PdfContainer from './PdfContainer';

class SalaryGenerator extends Component {
    constructor(props) {
        super(props);
        this.state = {
          id: '',
          salary: '10000',
          month: 'June'
        };
      }
    
      onChange = (event) => {
        const name = event.target.name;
        const value = event.target.value;
        this.setState((state) => {
          state[name] = value;
          return state;
        })
      }
      createPdf = (html) => Doc.createPdf(html);
    render() {
        return (
            <SiteWrapper>
            <Page.Card
                  title="Salary Slip Generator"
              ></Page.Card>
              <Grid.Col md={6} lg={8} className="align-self-center">
              <React.Fragment>
                    <section className="header-bar">
                    <span className="header">Generate Salary Slip</span>
                    </section>
                    <PdfContainer createPdf={this.createPdf}>
                    <React.Fragment>
                        <section className="flex-column">
                        <h2 className="flex">Employee ID</h2>
                        <section className="flex-row">
                            <input placeholder="id"
                            name="id"
                            value={this.state.id}
                            onChange={this.onChange} />
                        </section>
                        </section>
                    </React.Fragment>
                    </PdfContainer>
                </React.Fragment>
              </Grid.Col>
          </SiteWrapper>
        );
    }    
}

export default SalaryGenerator

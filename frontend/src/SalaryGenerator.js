import React, { Component } from 'react';
import './style.css';
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
    console.log(this.state);
    return (
      <React.Fragment>
        <section className="header-bar">
          <span className="header">Export React Component to PDF</span>
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
    );
  }
}

export default SalaryGenerator

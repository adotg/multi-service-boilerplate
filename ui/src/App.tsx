import React from "react";
import logo from "./logo.svg";
import "./App.css";

// function App() {}

class App extends React.Component<{}, { keyText: string; keyValue: string }> {
  constructor(props: React.Props<any>, ctx: React.Context<any>) {
    super(props, ctx);

    this.state = {
      keyText: "",
      keyValue: "",
    };
  }

  onSet() {
    console.log("set", this.state.keyText, this.state.keyValue);
  }

  onGet() {
    console.log("get", this.state.keyText);
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <p>Key</p>
          <input type="text" value={this.state.keyText} onChange={(e) => this.setState({ keyText: e.target.value })} />
          <p>value</p>
          <textarea value={this.state.keyValue} onChange={(e) => this.setState({ keyValue: e.target.value })} />
          <br />
          <button id="set_btn" onClick={this.onSet.bind(this)}>
            Set
          </button>
          <button id="get_btn" onClick={this.onGet.bind(this)}>
            Get
          </button>
        </header>
      </div>
    );
  }
}

export default App;

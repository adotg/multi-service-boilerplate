import React from "react";
import axios, { AxiosRequestConfig } from "axios";
import "./App.css";

// function App() {}

class App extends React.Component<{}, { keyText: string; keyValue: string; lastAccessTime: number }> {
  constructor(props: React.Props<any>, ctx: React.Context<any>) {
    super(props, ctx);

    this.state = {
      keyText: "",
      keyValue: "",
      lastAccessTime: -1,
    };
  }

  async onSet() {
    console.log("set", this.state.keyText, this.state.keyValue);
    if (this.state.keyText.trim() === "") {
      alert("key is mandatory");
      return;
    }

    if (this.state.keyValue.trim() === "") {
      alert("value is mandatory");
      return;
    }

    const resp = await axios({
      method: "post",
      // TODO take this from a variable?
      url: `http://localhost:8081/set/${this.state.keyText}`,
      data: this.state.keyValue,
      headers: {
        user_key: 123,
        "Content-Type": "text/plain",
        "X-Requested-With": "XMLHttpRequest",
      },
    });

    const { data } = resp;
    this.setState({ lastAccessTime: data.lastAccessTime, keyValue: data.data });
    console.log("Set value", data);
  }

  async onGet() {
    if (this.state.keyText.trim() === "") {
      alert("key is mandatory");
      return;
    }

    const resp = await axios({
      method: "get",
      // TODO take this from a variable?
      url: `http://localhost:8081/get/${this.state.keyText}`,
      headers: {
        user_key: 123,
        "Content-Type": "application/json",
        "X-Requested-With": "XMLHttpRequest",
      },
    });

    const { data } = resp;
    this.setState({ lastAccessTime: data.lastAccessTime, keyValue: data.data });
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
            Set (edited)
          </button>
          <button id="get_btn" onClick={this.onGet.bind(this)}>
            Get (edited)
          </button>
          <div>
            LastAccessTime: <span>{new Date(this.state.lastAccessTime * 1000).toString()}</span>
          </div>
        </header>
      </div>
    );
  }
}

export default App;

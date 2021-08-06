import React from "react";
import ReactDOM from "react-dom";
import { createBrowserHistory } from "history";
import { Router, Route, Switch } from "react-router-dom";
// import "assets/css/material-kit-react.css";
import ProblemPage from "views/ProblemPage/ProblemPage.js";
import HomePage from "views/HomePage/HomePage.js";
import ProblemSetPage from "views/ProblemSetPage/ProblemSetPage";
import NotFound from "views/NotFound";

var hist = createBrowserHistory();

ReactDOM.render(
  <Router history={hist}>
    <Switch>
      <Route exact path="/" component={HomePage} />
      <Route exact path="/problems/:problemName" component={ProblemPage} />
      <Route exact path="/problems" component={ProblemSetPage}/>
      <Route exact path='*'component={NotFound} />
    </Switch>
  </Router>,
  document.getElementById("root")
);
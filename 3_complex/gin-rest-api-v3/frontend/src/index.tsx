// TODO download modules ///////////////////////////////////////////////////////////////////////////////////////////////

import React from "react";
import { createRoot } from "react-dom/client";
import { Provider } from "react-redux";
import { store } from "./data/components/store";

// TODO custom modules /////////////////////////////////////////////////////////////////////////////////////////////////

import Router from "./data/components/router";

// TODO settings ///////////////////////////////////////////////////////////////////////////////////////////////////////

createRoot(document.getElementById("root")!).render(
  // <React.StrictMode>
  <Provider store={store}>
    <Router />
  </Provider>
  // </React.StrictMode>
);

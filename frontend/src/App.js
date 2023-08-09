import { Route, Routes } from "react-router-dom";

// components
import Navbar from "./components/Navbar";
import User from "./components/User";
import Home from "./components/Home";
import Authentication from "./components/Authentication";
import Order from "./components/Order";

// style
import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/authentication" element={<Authentication />} />
        <Route path="/user" element={<User />} />
        <Route path="/order" element={<Order />} />
      </Routes>
    </>
  );
}

export default App;

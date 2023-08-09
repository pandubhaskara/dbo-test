import { useState } from "react";
import { Link } from "react-router-dom";
import Icon from "./assets/hamburger.svg";
import "./styles/navbar.css";

const Layout = (props) => {
  const [isNavExpanded, setIsNavExpanded] = useState(false);

  return (
    <nav className="navigation">
      <a href="/" className="brand-name">
        DBO
      </a>
      <button
        className="hamburger"
        onClick={() => {
          setIsNavExpanded(!isNavExpanded);
        }}
      >
        <img src={Icon} alt="" />
      </button>
      <div
        className={
          isNavExpanded ? "navigation-menu expanded" : "navigation-menu"
        }
      >
        <ul>
          <li>
            <Link to={"/home"}>Home</Link>
          </li>
          <li>
            <Link to={"/user"}>User</Link>
          </li>
          <li>
            <Link to={"/order"}>Order</Link>
          </li>
        </ul>
        {props.children}
      </div>
    </nav>
  );
};
export default Layout;

import React from 'react';
import {Col,Container,Row} from 'react-bootstrap';
import { Link } from 'react-router-dom';

function Header() {
  return (
    <>
       <header style={styles.header}>
      <h2>My Website</h2>

      <nav>
        <ul style={styles.menu}>
            <li> <Link to="/">Home</Link> </li>
            <li> <Link to="/blog">Blog</Link> </li>
            <li> <Link to="/about">About Us</Link> </li>
            <li> <Link to="/contact">Contact Us</Link> </li>
        </ul>
      </nav>
    </header>
  </>
  );
}

const styles = {
  header: {
    display: "flex",
    justifyContent: "space-between",
    padding: "15px 30px",
    background: "#958181",
    color: "#fff",
  },
  menu: {
    display: "flex",
    gap: "20px",
    listStyle: "none",
  }
};

export default Header
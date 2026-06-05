import React from 'react'
import { Container , Row, Col} from 'react-bootstrap'

function Footer() {
  return (
    <footer
      style={{
        background: "#333",
        color: "#fff",
        textAlign: "center",
        padding: "15px",
        marginTop: "50px",
      }}
    >
      <p>© 2026 My Website. All Rights Reserved.</p>
    </footer>
  );
}

export default Footer


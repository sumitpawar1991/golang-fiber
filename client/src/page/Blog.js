import React, { useState, useEffect } from "react";
import { Col, Container, Row } from "react-bootstrap";
import axios from "axios";
import { Link } from "react-router-dom";

function Blog() {
  const [apiData, setApiData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const apiUrl = "http://localhost:8000";

      try {
        const response = await axios.get(apiUrl);

        if (
          response.status === 200 &&
          response.data.status === "ok"
        ) {
          setApiData(response.data.blog_records);
        }
      } catch (error) {
        console.log(error);
      }
    };

    fetchData();
  }, []);

  console.log(apiData);

  return (
    <Container>
      <Row>
        <Col xs={12} className="py-2">
          <h1 className="text-center">
            React application with Fiber backend
          </h1>
        </Col>

        {apiData.length > 0 ? (
          apiData.map((record, index) => (
            <Col key={index} xs={4} className="py-5 box">
              <div className="title">
                <Link to={`blog/${record.id}`}>{record.title}</Link>
                </div>
              <div>
                {record.post}
              </div>
            </Col>
          ))
        ) : (
          <Col xs={12}>
            <p>No data found</p>
          </Col>
        )}
      </Row>
    </Container>
  );
}

export default Blog;
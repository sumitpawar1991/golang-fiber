import React, { useState } from 'react'
import axios from 'axios'
import { Container,Row,Col,Form,Button } from 'react-bootstrap'
import { useNavigate } from 'react-router-dom'

function AddBlog() {
    const navigate = useNavigate();

    const [formData,setFormData] = useState({
        title:"",
        post:"",
    });

    const [errors,setErrors] = useState({});

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name] : e.target.value,
        })
    }

    const submitBlog = async (e) => {
        e.preventDefault();

        try {
                const response = await axios.post(
                    "http://localhost:8000/blog",
                    formData
                );

                if (response.data.status === "ok") {
                    navigate("/");
                }
                } catch (error) {
                if (error.response?.data?.errors) {
                    setErrors(error.response.data.errors);
                }
                }
    };

   return (
    <Container>
      <Row className="justify-content-center">
        <Col md={6}>
          <h2 className="my-4">Add New Blog</h2>

          <Form onSubmit={submitBlog}>
            <Form.Group className="mb-3">
              <Form.Label>Title</Form.Label>
              <Form.Control
                type="text"
                name="title"
                value={formData.title}
                onChange={handleChange}
              />
              <small className="text-danger">
                {errors.title}
              </small>
            </Form.Group>

            <Form.Group className="mb-3">
              <Form.Label>Post</Form.Label>
              <Form.Control
                as="textarea"
                rows={5}
                name="post"
                value={formData.post}
                onChange={handleChange}
              />
              <small className="text-danger">
                {errors.post}
              </small>
            </Form.Group>

            <Button type="submit">
              Save Blog
            </Button>
          </Form>
        </Col>
      </Row>
    </Container>
  );
}

export default AddBlog;
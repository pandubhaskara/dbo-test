import { Form, Table, Button, Modal } from "react-bootstrap";
import { useState, createRef } from "react";
import user from "../../data/user.json";

export default function User() {
  const [users, setUser] = useState(user);
  const [status, setStatus] = useState("");
  const [message, setMessage] = useState("");
  const [showModal, setShowModal] = useState(false);

  const formData = createRef();

  const add = (event) => {
    event.preventDefault();

    if (
      !formData.current.name.value ||
      !formData.current.email.value ||
      !formData.current.mobile_number.value
    ) {
      setStatus("Error");
      setMessage("Failed add user");
      setShowModal(true);
      return;
    }

    const newUser = {
      name: formData.current.name.value,
      email: formData.current.email.value,
      mobile_number: formData.current.mobile_number.value,
    };
    
    setUser([...users, newUser]);
    setStatus("Success");
    setMessage("Successfully add user");
    setShowModal(true);
  };

  const handleClose = () => setShowModal(false);

  return (
    <div>
      <Modal show={showModal} onHide={handleClose}>
        <Modal.Header closeButton>
          <Modal.Title>{status}</Modal.Title>
        </Modal.Header>
        <Modal.Body>{message}</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>

      <div
        style={{
          alignItems: "center",
          justifyContent: "center",
          padding: "20px",
        }}
      >
        <Form onSubmit={add} ref={formData}>
          <Form.Group>
            <Form.Label>Name:</Form.Label>
            <Form.Control type="text" placeholder="John Doe" name="name" />
          </Form.Group>

          <Form.Group>
            <Form.Label>Email:</Form.Label>
            <Form.Control
              type="email"
              placeholder="johndoe@mail.com"
              name="email"
            />
          </Form.Group>

          <Form.Group>
            <Form.Label>Mobile Number:</Form.Label>
            <Form.Control
              type="text"
              placeholder="+628123456789"
              name="mobile_number"
            />
          </Form.Group>

          <br />

          <Button variant="primary" type="submit">
            Add New User
          </Button>
        </Form>
      </div>

      <Table striped bordered hover responsive variant="dark">
        <thead>
          <tr>
            <th>No</th>
            <th>Name</th>
            <th>Email</th>
            <th>Mobile Number</th>
          </tr>
        </thead>
        <tbody>
          {users.map((item, index) => {
            return (
              <tr key={index}>
                <td>{index + 1}</td>
                <td>{item.name}</td>
                <td>{item.email}</td>
                <td>{item.mobile_number}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </div>
  );
}

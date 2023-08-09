import { Form, Table, Button, Modal, Row, Col } from "react-bootstrap";
import { useState, createRef } from "react";
import orders from "../../data/order.json";

export default function Order() {
  const [order, setOrder] = useState(orders);
  const [status, setStatus] = useState("");
  const [message, setMessage] = useState("");
  const [showModal, setShowModal] = useState(false);

  const formData = createRef();

  const add = (event) => {
    event.preventDefault();

    const newOrder = {
      user_id: formData.current.user_id?.value,
      product_id: formData.current.product_id?.value,
      name: formData.current.name?.value,
      price: formData.current.price?.value,
      quantity: formData.current.quantity?.value,
      total: formData.current.total?.value,
      supplier_name: formData.current.supplier_name?.value,
    };

    setOrder([...order, newOrder]);
    setStatus("Success");
    setMessage("Successfully add order");
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
            <Row>
              <Col lg={6} md={6} sm={12} xs={12}>
                <Form.Label>User ID:</Form.Label>
                <Form.Control
                  required="true"
                  type="number"
                  placeholder="1"
                  name="user_id"
                />
                <Form.Label>Product ID:</Form.Label>
                <Form.Control
                  required="true"
                  type="number"
                  placeholder="1"
                  name="product_id"
                />
                <Form.Label>Name:</Form.Label>
                <Form.Control
                  required="true"
                  type="text"
                  placeholder="Pipa"
                  name="name"
                />
              </Col>

              <Col lg={6} md={6} sm={12} xs={12}>
                <Form.Label>Price:</Form.Label>
                <Form.Control
                  required="true"
                  type="number"
                  placeholder="10000"
                  name="price"
                />
                <Form.Label>Quantity:</Form.Label>
                <Form.Control
                  required="true"
                  type="number"
                  placeholder="1"
                  name="quantity"
                />
                <Form.Label>Supplier:</Form.Label>
                <Form.Control
                  required="true"
                  as="select"
                  custom
                  name="supplier_name"
                >
                  <option key="blankChoice" hidden value="">
                    Select
                  </option>
                  <option value="Djabesmen">Djabesmen</option>
                  <option value="Rucika">Rucika</option>
                </Form.Control>
              </Col>
            </Row>
          </Form.Group>

          <br />

          <Button variant="primary" type="submit">
            Add New Order
          </Button>
        </Form>
      </div>

      <Table striped bordered hover responsive variant="dark">
        <thead>
          <tr>
            <th>No</th>
            <th>User ID</th>
            <th>Product ID</th>
            <th>Product Name</th>
            <th>Product Price</th>
            <th>Quantity</th>
            <th>Total</th>
            <th>Supplier</th>
          </tr>
        </thead>
        <tbody>
          {order.map((item, index) => {
            return (
              <tr key={index}>
                <td>{index + 1}</td>
                <td>{item.user_id}</td>
                <td>{item.product_id}</td>
                <td>{item.name}</td>
                <td>{item.price}</td>
                <td>{item.quantity}</td>
                <td>
                  {new Intl.NumberFormat("id-ID", {
                    style: "currency",
                    currency: "IDR",
                  }).format(item.price * item.quantity)}
                </td>
                <td>{item.supplier_name}</td>
              </tr>
            );
          })}
        </tbody>
      </Table>
    </div>
  );
}

import { AspectRatio, Button, Container, createStyles, Group, Space, Tooltip } from "@mantine/core";
import axios from "axios";
import { useState } from "react";
import { ArrowForwardUp, Bike, Car, PlaneInflight, Train, Walk } from 'tabler-icons-react';

const useStyles = createStyles((theme) => ({
  button: {
    borderRadius: 0,

    '&:not(:first-of-type)': {
      borderLeftWidth: 0,
    },

    '&:first-of-type': {
      borderTopLeftRadius: theme.radius.sm,
      borderBottomLeftRadius: theme.radius.sm,
    },

    '&:last-of-type': {
      borderTopRightRadius: theme.radius.sm,
      borderBottomRightRadius: theme.radius.sm,
    },
  },
}));

export function Eat() {

  const [city, setCity] = useState('')
  const [radius, setRadius] = useState('')


  const retrieveSleep = (event: any) => {
    axios.defaults.withCredentials = true
    event.preventDefault()
    let params = {
      city: city,
      constraints: {
        radius: 10000,
      }
    };
    console.log("event: ", params);

    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/eat',
      data: params,
    })
      .then((response) => {
        console.log(response.data);
        return response.data;
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  return (
    <form onSubmit={retrieveSleep}>
      <h1 className="h3 mb-3 fw-normal">Eat Activities</h1>

      <input type="text" className="form-control" placeholder="City" required
        onChange={e => setCity(e.target.value)}
      />

      <input type="text" className="form-control" placeholder="Radius" required
        onChange={e => setRadius(e.target.value)}
      />

      <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
    </form>
    // <Container size={720}>
    //   <Group grow spacing={0}>
    //     <Button variant="default" className={classes.button}>
    //       <ArrowForwardUp />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Car />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Train />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Walk />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <Bike />
    //     </Button>
    //     <Button variant="default" className={classes.button}>
    //       <PlaneInflight />
    //     </Button>
    //   </Group>

    //   <Space h="xl" />

    //   <AspectRatio ratio={16 / 9}>
    //     <iframe
    //       src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d10500.902039411167!2d2.2913514895690534!3d48.85391001859108!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x47e66e2964e34e2d%3A0x8ddca9ee380ef7e0!2sEiffel%20Tower!5e0!3m2!1sen!2sru!4v1653233639984!5m2!1sen!2sru"
    //       title="Google map"
    //       frameBorder="0"
    //     />
    //   </AspectRatio>
    // </Container>
  )
}
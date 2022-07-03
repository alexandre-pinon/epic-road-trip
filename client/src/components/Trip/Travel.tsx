import {
  AspectRatio,
  Button,
  Card,
  Container,
  createStyles,
  Group,
  Space,
  Tooltip,
  Text,
  SimpleGrid,
  Paper,
  Grid
} from "@mantine/core";
import { ArrowForwardUp, Bike, Car, PlaneInflight, Train, Walk } from 'tabler-icons-react';
import axios from "axios";
import {SetStateAction, useEffect, useState} from "react";

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


export function Travel(props: any) {
  const {classes} = useStyles();

  const [cityDeparture, setCityDeparture] = useState('')
  const [cityArrival, setCityArrival] = useState('')
  const [duration, setDuration] = useState('')
  const [startDate, setStartDate] = useState('')
  const [endDate, setEndDate] = useState('')
  const [price, setPrice] = useState('')


  const [train, setTrain] = useState([{
    cityDeparture,
    cityArrival,
    duration,
    startDate,
    endDate
  }
  ])
  const [plane, setPlane] = useState([{
    cityDeparture,
    cityArrival,
    duration,
    startDate,
    endDate,
    price
  }])


  const NON = () => {
    let maVoiture = {
      make: 'Ford',
      model: 'Mustang',
      year: 1969
    };
    let maVoiture2 = {
      make: 'Ford',
      model: 'Mustang',
      year: 1969
    };
    let myList = [];
    myList.push(maVoiture)

    // @ts-ignore
    setTrain(myList)
    console.log(train)
  }




  const planeTravel = () => {
    let params = {
      adults: 1,
      departureDate: "2022-08-08T15:04:05Z",
      destinationLocation: "Paris",
      maxPrice: 10000,
      originLocation: "London",
    };
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/travel/air',
      data: params
    }).then(res => {
      let myList2: any = [];
      res.data.data.forEach((data: any ) => {
        let travelInfo = {
          cityDeparture: data.arrival.city,
          cityArrival: data.departure.city,
          duration: data.duration,
          startDate: data.startdate,
          endDate: data.enddate,
          price: data.price
        }
        myList2.push(travelInfo)
      })
      setPlane(myList2)
      console.log(myList2)
    });
  };

  const trainTravel = () => {
    let params = {
      departureTime: "2022-08-18T15:04:05Z",
      destination: "Paris",
      origin: "London",
    };
    axios({
      method: 'post',
      url: 'http://localhost:8000/api/v1/roadtrip/travel/ground',
      data: params
    }).then(res => {
      let myList2: any = [];
      res.data.data.forEach((data: any ) => {
        let travelInfo = {
          cityDeparture: data.arrival.city,
          cityArrival: data.departure.city,
          duration: data.duration,
          startDate: data.startdate,
          endDate: data.enddate,
        }
        myList2.push(travelInfo)
      })
      setTrain(myList2)
      console.log(myList2)
    });
  };

  useEffect(() => {
    planeTravel();
    trainTravel();
  }, []);



  return (
      <Container size={1000}>
        <Group grow spacing={0} position="apart">

          <Group grow spacing={0} position="apart">
            <SimpleGrid cols={1}>
              <>
                <Button onClick={planeTravel} variant="default" className={classes.button}>
                  <PlaneInflight />
                </Button>

                <ul>
                  {
                    plane.length ? (
                            plane.map((item, index) => (
                                <Paper shadow="xl" p="md" withBorder key={index}>
                                  <Grid><Text weight={700}>Departure :  </Text> <Text> -  {item.cityDeparture}</Text></Grid>
                                  <Grid><Text weight={700}>Arrival :  </Text> <Text> -  {item.cityArrival}</Text></Grid>
                                  <Grid><Text weight={700}>Start Date :  </Text> <Text> -  {item.startDate}</Text></Grid>
                                  <Grid><Text weight={700}>End Date :  </Text> <Text> -  {item.endDate}</Text></Grid>
                                  <Grid><Text weight={700}>Duration :  </Text> <Text> -  {item.duration}</Text></Grid>
                                  <Grid><Text weight={700}>Price :  </Text> <Text>{ - item.price}</Text></Grid>
                                </Paper>
                            ))
                        )
                        : null
                  }
                </ul>
              </>
            </SimpleGrid>

          </Group>

          <Group grow spacing={0} position="apart">
            <SimpleGrid cols={1}>
              <>
                <Button variant="default" className={classes.button}>
                  <Train />
                </Button>
                <ul>
                  {
                    train.length ? (
                            train.map((item, index) => (
                                <Paper shadow="xl" p="md" withBorder key={index}>
                                  <Grid><Text weight={700}>Departure :  </Text> <Text> -  {item.cityDeparture}</Text></Grid>
                                  <Grid><Text weight={700}>Arrival :  </Text> <Text> -  {item.cityArrival}</Text></Grid>
                                  <Grid><Text weight={700}>Start Date :  </Text> <Text> -  {item.startDate}</Text></Grid>
                                  <Grid><Text weight={700}>End Date :  </Text> <Text> -  {item.endDate}</Text></Grid>
                                  <Grid><Text weight={700}>Duration :  </Text> <Text> -  {item.duration}</Text></Grid>
                                </Paper>
                            ))
                        )
                        : null
                  }
                </ul>
              </>
            </SimpleGrid>
          </Group>
        </Group>









        <Space h="xl" />

      {/*
      <AspectRatio ratio={16 / 9}>
        <iframe
          src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d10500.902039411167!2d2.2913514895690534!3d48.85391001859108!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x47e66e2964e34e2d%3A0x8ddca9ee380ef7e0!2sEiffel%20Tower!5e0!3m2!1sen!2sru!4v1653233639984!5m2!1sen!2sru"
          title="Google map"
          frameBorder="0"
        />
      </AspectRatio>
      */}
    </Container>
  )
}
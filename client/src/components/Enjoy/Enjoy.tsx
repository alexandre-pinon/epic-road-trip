import { AspectRatio, Avatar, Button, Card, Container, createStyles, Group, Space, Tooltip, Text, Image, Paper, Grid, Center } from "@mantine/core";
import axios from "axios";
import { useState } from "react";
import { ArrowForwardUp, Bike, Car, PlaneInflight, Search, Train, Walk } from 'tabler-icons-react';
import { useNavigate } from "react-router-dom";

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


export function Enjoy({ fulTrip }: any) {
  const navigate = useNavigate();

  const goTravel = async () => {
    console.log("Go to travel page!")
    navigate('/travel');
  };

  const goSleep = async () => {
    console.log("Go to sleep page!")
    navigate('/sleep');
  };

  const [id, setId] = useState(0)
  const [city, setCity] = useState('')
  const [radius, setRadius] = useState('')
  const [name, setName] = useState('')
  const [rating, setRating] = useState('')
  const [vicinity, setVicinity] = useState('')

  const [enjoy, setEnjoy] = useState([{
    id,
    name,
    rating,
    vicinity
  }])

  let [selectedEnjoy, setSelectedEnjoy] = useState('')


  const [toggleEnjoy, setToggleEnjoy] = useState(false)

  const retrieveEnjoy = (event: any) => {
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
      url: 'http://localhost:8000/api/v1/roadtrip/enjoy',
      data: params,
    })
      .then((response) => {
        console.log(response.data);
        let id = 0
        let enjoyActivities: any = [];
        response.data.data.forEach((data: any) => {
          let activities = {
            id: id,
            name: data.name,
            rating: data.rating,
            vicinity: data.vicinity,
          }
          enjoyActivities.push(activities)
          id++;
        })
        setEnjoy(enjoyActivities)
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const showEnjoy = () => {
    setToggleEnjoy(!toggleEnjoy)
  }

  const saveEnjoy = () => {
    console.log(fulTrip)
  }

  const selectEnjoy = (id: number, type: string) => {

    setSelectedEnjoy(type)
    console.log(id)
    setId(id)

    if (type == "Enjoy") {
      fulTrip.setEnjoyName(enjoy[id])
    }

    console.log("fulTrip content: ", fulTrip)
  }

  return (
    <Container>
      <form onSubmit={retrieveEnjoy}>
        <h1 className="h3 mb-3 fw-normal">Enjoy Activities</h1>

        <input type="text" className="form-control" placeholder="City" required
          onChange={e => setCity(e.target.value)}
        />

        <input type="text" className="form-control" placeholder="Radius" required
          onChange={e => setRadius(e.target.value)}
        />

        <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
      </form>

      <Button onClick={showEnjoy} variant="default">
        <PlaneInflight />
      </Button>

      <ul>
        {
          toggleEnjoy ? (
            enjoy.map((item) => (
              <Paper shadow="xl" p="md" withBorder key={item.id}>
                <Grid><Text weight={700}>Name :  </Text> <Text> -  {item.name}</Text></Grid>
                <Grid><Text weight={700}>Rating :  </Text> <Text> -  {item.rating}</Text></Grid>
                <Grid><Text weight={700}>Vicinity :  </Text> <Text> -  {item.vicinity}</Text></Grid>
                <Center><Button onClick={() => selectEnjoy(item.id, 'Enjoy')} >Select this enjoy </Button></Center>
              </Paper>
            ))
          )
            : null
        }
      </ul>

      <div>
        {
          selectedEnjoy === "Enjoy" ? (
            <>
              <Center><h3>SELECTED TRAVEL : </h3></Center>
              <Paper shadow="xl" p="md" withBorder >
                <Grid><Text weight={700}>Enjoy ID  {enjoy[id].id}</Text> </Grid>
                <Grid><Text weight={700}>Name :  </Text> <Text> -  {enjoy[id].name}</Text></Grid>
                <Grid><Text weight={700}>Rating :  </Text> <Text> -  {enjoy[id].rating}</Text></Grid>
                <Grid><Text weight={700}>Vicinity :  </Text> <Text> -  {enjoy[id].vicinity}</Text></Grid>
                <Center><Button onClick={() => selectEnjoy(enjoy[id].id, 'Enjoy')} >Confirm this enjoy </Button></Center>
              </Paper>
            </>
          )
            : null
        }
      </div>

      <Center>
        <Button onClick={goTravel} rightIcon={<Search size={18} />} variant="light" radius="xl">
          Go back
        </Button>
        <Button onClick={goSleep} rightIcon={<Search size={18} />} variant="light" radius="xl">
          Search for Activities
        </Button>
      </Center>
    </Container>


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

function enjoyActivities(enjoyActivities: any) {
  throw new Error("Function not implemented.");
}


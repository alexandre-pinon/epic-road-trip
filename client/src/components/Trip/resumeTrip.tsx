import { useNavigate } from 'react-router-dom';
import {Avatar, Button, Center, Title, Card, Image, Text, Paper, Space, Grid} from '@mantine/core';
import { useForm } from '@mantine/form';
import React, {useState} from "react";





export function ResumeTrip({fulTrip} : any) {
    const navigate = useNavigate();
    const [subTrip, setSubTrip] = useState(
        [
            {
                "startCity":"Montréal",
                "endCity":"Toronto",
                "Trajets":"Bus",
                "Activities": [
                    {
                        "hotel":"Dans les locaux d'AB Tasty",
                        "restaurant":"La cafet d'AB Tasty",
                        "Entertainment1":"Coder des tests Cypress",
                        "Entertainment2":"Faire des tickets sur JIRA"
                    }
                ]
            },
            {
                "startCity":"Tokyo",
                "endCity":"Kyoto",
                "Trajets":"Trains",
                "Activities": [
                    {
                        "hotel":"Chez Eichiro Oda",
                        "restaurant":"Le Baratié",
                        "Entertainment1":"Trouver le One Piece",
                        "Entertainment2":"Libérer Wano"
                    }
                ]
            },
            {
                "startCity":"Delhi",
                "endCity":"Bombay",
                "Trajets":"Avions",
                "Activities": [
                    {
                        "hotel":"Chez Tharick",
                        "restaurant":"Saravanha",
                        "Entertainment1":"Voir le Taj Mahal",
                        "Entertainment2":"Rendre visite à la famille de Tharick"
                    }
                ]
            }
        ])

    console.log(fulTrip)
    console.log(fulTrip.startDateValue.toDateString())



    const goToHome = async () => {
        console.log("Going back home!")
        navigate('/');
    };

    const addSubTrip = async () => {
        console.log("Add a sub-trip")
        navigate('/');
    };

    return (
        <>
            <Center><Title order={1}>This is the summary of your Trip</Title></Center>

            <Space h="xl" />

            <Paper shadow="xs" radius="lg" p="lg" withBorder>
                <Center><h2>Your TRIP</h2></Center>
                <h3>Cities</h3>
                <Grid><Text weight={700}>Start City :  </Text><Text> &nbsp;{fulTrip.startCity}</Text></Grid>
                <Grid><Text weight={700}>End City :  </Text><Text> &nbsp;{fulTrip.endCity}</Text></Grid>
                <Space h="md" />
                <h3>Dates</h3>
                <Grid><Text weight={700}>Start Date :  </Text><Text> &nbsp;{fulTrip.startDateValue.toDateString()}</Text></Grid>
                <Grid><Text weight={700}>End Date :  </Text><Text> &nbsp;{fulTrip.endDateValue.toDateString()}</Text></Grid>
                <Space h="md" />
                <h3>Itinerary</h3>
                <Grid><Text weight={700}>Departure city :  </Text><Text> &nbsp;{fulTrip.selectedTravel.cityDeparture}</Text></Grid>
                <Grid><Text weight={700}>Arrival city :  </Text><Text> &nbsp;{fulTrip.selectedTravel.cityArrival}</Text></Grid>
                <Grid><Text weight={700}>Departure time :  </Text><Text> &nbsp;{fulTrip.selectedTravel.startDate}</Text></Grid>
                <Grid><Text weight={700}>Arrival time :  </Text><Text> &nbsp;{fulTrip.selectedTravel.endDate}</Text></Grid>
                <Grid><Text weight={700}>Duration :  </Text><Text> &nbsp;{fulTrip.selectedTravel.duration}</Text></Grid>
                <Space h="md" />
                <h3>Enjoy Activities</h3>
                <Grid><Text weight={700}>Name :  </Text><Text> &nbsp;{fulTrip.enjoyName.name}</Text></Grid>
                <Grid><Text weight={700}>Vicinity :  </Text><Text> &nbsp;{fulTrip.enjoyName.vicinity}</Text></Grid>
                <Grid><Text weight={700}>Rating :  </Text><Text> &nbsp;{fulTrip.enjoyName.rating}</Text></Grid>
            </Paper>

            <Paper shadow="xs" radius="lg" p="lg" withBorder>
                {subTrip.map((subTrip, index) => (
                    <>
                        <Center><h2>Your sub-trip {index}</h2></Center>
                        <Paper shadow="md" radius="lg" p="lg" withBorder>
                            <Text>Start city : {subTrip.startCity}</Text>
                            <Text>End city : {subTrip.endCity}</Text>
                            <Text>Your mode of transport : {subTrip.Trajets}</Text>
                            <Text>{subTrip.Activities[0].hotel}</Text>
                            <Text>{subTrip.Activities[0].restaurant}</Text>
                            <Text>{subTrip.Activities[0].Entertainment1}</Text>
                            <Text>{subTrip.Activities[0].Entertainment1}</Text>

                            {subTrip.Activities.map((activities) => (
                                <>
                                    <Text>{activities.hotel}</Text>
                                    <Text>{activities.restaurant}</Text>
                                    <Text>{activities.Entertainment1}</Text>
                                    <Text>{activities.Entertainment2}</Text>
                                </>

                            ))}


                        </Paper>
                        <h1></h1>
                    </>
                ))}

            </Paper>

            <Space h="xl" />

            <Center>
                <Button onClick={goToHome} variant="light" radius="xl">
                    Retour
                </Button>
                <Button onClick={addSubTrip} variant="light" radius="xl">
                    Add an other sub-Trip
                </Button>
            </Center>
        </>
    )
}
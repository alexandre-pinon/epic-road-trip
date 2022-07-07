import { useNavigate } from 'react-router-dom';

import React, { useState } from "react";
import { ActionIcon, Button, Center, Container, createStyles, Modal, Space, TextInput, Title, Autocomplete } from "@mantine/core";
import { BuildingSkyscraper, CalendarEvent, Search, BuildingCommunity, CalendarMinus, ArrowBackUp, Plane } from "tabler-icons-react";
import { Calendar, DatePicker } from "@mantine/dates";
import { props } from "kea";
import cityObjects from '../../data/iata_codes.json';



export function StartEndTrip({ fulTrip }: any) {
    const navigate = useNavigate();

    const goToHome = async () => {
        console.log("Going back home!")
        navigate('/');
    };

    const goTravel = async () => {
        console.log("Go to travel page!")
        navigate('/travel');
    };

    const city = Object.keys(cityObjects)

    return (
        <>
            <Center>
                <Title style={{ color: "#616161 " }} order={2} data-testid="title"> Start & End of your sub-trip </Title>
            </Center>

            <Space h="xl" />

            <Center>
                <Title style={{ color: "#616161 " }} order={4} data-testid="cityTitle">Choose the city of departure & arrival of your sub-trip </Title>
            </Center>

            <Space h="lg" />


            <Container size={540}>
                {/* Ville de départ */}
                <Autocomplete
                    icon={<BuildingCommunity size={18} />}
                    radius="xl"
                    size="md"
                    placeholder="Start city of your Trip"
                    rightSectionWidth={42}
                    {...props}
                    data={city}
                    value={fulTrip.startCity}
                    onChange={(event) => fulTrip.setStartCity(event)}
                    data-testid="StartCity"
                />

                <Space h="xl" />

                {/* Ville d'arrivée */}

                <Autocomplete
                    icon={<BuildingCommunity size={18} />}
                    radius="xl"
                    size="md"
                    placeholder="End city of your Trip"
                    rightSectionWidth={42}
                    {...props}
                    data={city}
                    onChange={(event) => fulTrip.setEndCity(event)}
                    data-testid="endCity"
                />

            </Container>
            <Space h="xl" />

            <Center>
                <Title style={{ color: "#616161 " }} order={4} data-testid="dateTitle">Choose the date of departure & arrival of your sub-trip </Title>
            </Center>

            <Space h="xl" />

            <Container size={540}>

                <DatePicker
                    icon={<CalendarEvent size={18} />}
                    placeholder="Start date of your Trip"
                    radius="xl"
                    onChange={fulTrip.setStartDate}
                    data-testid="startDate"
                />

                <Space h="xl" />

                <DatePicker
                    icon={<CalendarMinus size={18} />}
                    placeholder="End date of your Trip"
                    radius="xl"
                    onChange={fulTrip.setEndDate}
                    data-testid="endDate"
                />

                <Space h="xl" />

                <Center>
                    <Button onClick={goToHome} rightIcon={<ArrowBackUp size={18} />} compact variant="subtle" radius="xs" data-testid="goBack">
                        Go back
                    </Button>
                    <Button onClick={goTravel} rightIcon={<Plane size={18} />} compact variant="subtle" radius="xs" data-testid="nextStep">
                        Search for itineraries
                    </Button>
                </Center>
            </Container>
        </>
    )
}
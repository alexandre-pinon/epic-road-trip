import { useNavigate } from 'react-router-dom';

import React, {useState} from "react";
import {ActionIcon, Button, Center, Container, createStyles, Modal, Space, TextInput, Title, Autocomplete} from "@mantine/core";
import {BuildingSkyscraper, CalendarEvent, Search, BuildingCommunity, CalendarMinus} from "tabler-icons-react";
import {Calendar, DatePicker} from "@mantine/dates";
import {props} from "kea";
import cityObjects from '../../data/iata_codes.json';



export function StartEndTrip( {fulTrip} : any) {
    const navigate = useNavigate();

    const goToHome = async () => {
        console.log("Going back home!")
        navigate('/');
    };

    const city = Object.keys(cityObjects)

    return (
        <>
            <Center>
                <Title order={2}> Start & End of your sub-trip </Title>
            </Center>

            <Space h="xl" />

            <Center>
                <Title order={3}> Please choose the city of departure and arrival of your sub-trip </Title>
            </Center>

            <Space h="lg" />


            <Container size={540}>
                {/* Ville de départ */}
                <Autocomplete
                    icon={<BuildingCommunity size={18} />}
                    radius="xl"
                    size="md"
                    placeholder="Ville de départ"
                    rightSectionWidth={42}
                    {...props}
                    data={city}
                    value={fulTrip.startCity}
                    onChange={(event) => fulTrip.setStartCity(event)}
                />

                <Space h="xl" />

                {/* Ville d'arrivée */}

                <Autocomplete
                    icon={<BuildingCommunity size={18} />}
                    radius="xl"
                    size="md"
                    placeholder="Ville d'arrivée"
                    rightSectionWidth={42}
                    {...props}
                    data={city}
                    onChange={(event) => fulTrip.setEndCity(event)}
                />

            </Container>
            <Space h="xl" />

            <Center>
                <Title order={3}> Please choose the date of departure and arrival of your sub-trip </Title>
            </Center>

            <Space h="xl" />

            <Container size={540}>

                <DatePicker
                    icon={<CalendarEvent size={18} />}
                    placeholder="Start date of your Trip"
                    radius="xl"
                    onChange={fulTrip.setStartDate}
                />

                <Space h="xl" />

                <DatePicker
                    icon={<CalendarMinus size={18} />}
                    placeholder="End date of your Trip"
                    radius="xl"
                    onChange={fulTrip.setEndDate}
                />

                <Space h="xl" />

                <Center>
                    <Button rightIcon={<Search size={18} />} variant="light" radius="xl">
                        Search for itineraries
                    </Button>
                </Center>
            </Container>
        </>
    )
}
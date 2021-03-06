import React, { useState } from 'react'
import { createStyles, Container, UnstyledButton, Text, Card, SimpleGrid, Space, Title, Grid, Button, Center } from '@mantine/core';

import {
  HotelService,
  BrandAirbnb,
  Run,
  BrandTripadvisor,
  BrandBooking,
  Plus,
} from 'tabler-icons-react';
import { useNavigate } from 'react-router-dom';

const mockdata = [
  { title: 'Hotels', icon: HotelService, color: 'dark' },
  { title: 'Vacation Rentals', icon: BrandAirbnb, color: 'dark' },
  { title: 'Things To Do', icon: Run, color: 'dark' },
  { title: 'Restaurants', icon: BrandTripadvisor, color: 'dark' },
  { title: 'Travel Forums', icon: BrandBooking, color: 'dark' },
  { title: 'More ...', icon: Plus, color: 'dark' },
];

const useStyles = createStyles((theme) => ({
  card: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.gray[0],
  },

  title: {
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
    fontWeight: 700,
  },

  item: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: 'center',
    borderRadius: theme.radius.md,
    height: 90,
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
    transition: 'box-shadow 150ms ease, transform 100ms ease',

    '&:hover': {
      boxShadow: `${theme.shadows.md} !important`,
      transform: 'scale(1.05)',
    },
  },

  // [START] Hero with Background Image
  wrapper: {
    position: 'relative',
    paddingTop: 180,
    paddingBottom: 130,
    backgroundImage:
      'url(https://images2.alphacoders.com/106/1062364.png)',
    backgroundSize: 'cover',
    backgroundPosition: 'center',

    '@media (max-width: 520px)': {
      paddingTop: 80,
      paddingBottom: 50,
    },
  },

  inner: {
    position: 'relative',
    zIndex: 1,
  },
  // [END] Hero with Background Image

  // [START] Card with icon features
  card2: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
  },

  imageSection: {
    padding: theme.spacing.md,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    borderBottom: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
      }`,
  },

  label: {
    marginBottom: theme.spacing.xs,
    lineHeight: 1,
    fontWeight: 700,
    fontSize: theme.fontSizes.xs,
    letterSpacing: -0.25,
    textTransform: 'uppercase',
  },

  section: {
    padding: theme.spacing.md,
    borderTop: `1px solid ${theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
      }`,
  },

  icon: {
    marginRight: 5,
    color: theme.colorScheme === 'dark' ? theme.colors.dark[2] : theme.colors.gray[5],
  },
  // [END] Card with icon features
}));

export function City() {
  const { classes, theme, cx } = useStyles();
  const navigate = useNavigate();

  const goToHome = async () => {
    console.log("Going back home!")
    navigate('/');
  };

  const items = mockdata.map((item) => (
    <UnstyledButton key={item.title} className={classes.item}>
      <item.icon color={theme.colors[item.color][6]} size={32} />
      <Text size="xs" mt={7}>
        {item.title}
      </Text>
    </UnstyledButton>
  ));

  return (
    <Container>
      <Title order={2}>Explorez "Ville de d??part"</Title>

      <Space h="md" />

      <Card withBorder radius="md" className={classes.card}>
        <SimpleGrid cols={6} mt="md">
          {items}
        </SimpleGrid>
      </Card>

      <Space h="xl" />
      <Space h="xl" />
      <Space h="xl" />

      <Title order={2}>Suggestions</Title>

      <Space h="xl" />
      <Space h="xl" />

      <Grid>
        <Grid.Col md={6} lg={3}>
          <Title order={5}>Se divertir</Title>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            Activit?? 1
          </Button>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            Activit?? 2
          </Button>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            Activit?? 3
          </Button>
        </Grid.Col>
      </Grid>

      <Space h="xl" />
      <Space h="xl" />

      <Grid>
        <Grid.Col md={6} lg={3}>
          <Title order={5}>Se reposer</Title>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            H??tel 1
          </Button>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            H??tel 2
          </Button>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            H??tel 3
          </Button>
        </Grid.Col>
      </Grid>

      <Space h="xl" />
      <Space h="xl" />

      <Grid>
        <Grid.Col md={6} lg={3}>
          <Title order={5}>Se restaurer</Title>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            Restaurant 1
          </Button>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            Restaurant 2
          </Button>
        </Grid.Col>
        <Grid.Col md={6} lg={3}>
          <Button variant="light" color="blue">
            Restaurant 3
          </Button>
        </Grid.Col>
      </Grid>

      <Space h="xl" />
      <Space h="xl" />

      <Center>
        <Button onClick={goToHome} variant="light" radius="xl">
          Retour
        </Button>
      </Center>

    </Container >
  );

}
describe('The resume trip', () => {

    it('Check text on the ResumeTrip page', () => {

        cy.visit('http://localhost:3000/resumeTrip/')

        cy.get('[data-testid="Big-title"]')
            .contains('This is your resume trip 🚌');

        cy.get('[data-testid="yourTrip"]')
            .contains('Your trip');

        cy.get(':nth-child(4) > .mantine-dj2qqh').contains('Start City :')
        cy.get(':nth-child(5) > .mantine-dj2qqh').contains('End City :')

        cy.get(':nth-child(9) > .mantine-dj2qqh').contains('Start Date :')
        cy.get(':nth-child(10) > .mantine-dj2qqh').contains('End Date :')

        cy.get(':nth-child(14) > .mantine-dj2qqh').contains('Departure city :')
        cy.get(':nth-child(15) > .mantine-dj2qqh').contains('Arrival city :')
        cy.get(':nth-child(16) > .mantine-dj2qqh').contains('Departure time :')
        cy.get(':nth-child(17) > .mantine-dj2qqh').contains('Arrival time :')
        cy.get(':nth-child(18) > .mantine-dj2qqh').contains('Duration :')

        cy.get(':nth-child(25) > .mantine-dj2qqh').contains('Name :')
        cy.get(':nth-child(26) > .mantine-dj2qqh').contains('Vicinity :')
        cy.get(':nth-child(27) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(25) > .mantine-dj2qqh').contains('Name :')
        cy.get(':nth-child(26) > .mantine-dj2qqh').contains('Vicinity :')
        cy.get(':nth-child(27) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(31) > .mantine-dj2qqh').contains('Restaurant name :')
        cy.get(':nth-child(32) > .mantine-dj2qqh').contains('Address :')
        cy.get(':nth-child(33) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(37) > .mantine-dj2qqh').contains('Restaurant name :')
        cy.get(':nth-child(38) > .mantine-dj2qqh').contains('Address :')
        cy.get(':nth-child(39) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(43) > .mantine-dj2qqh').contains('Hotel :')
        cy.get(':nth-child(44) > .mantine-dj2qqh').contains('Address :')
        cy.get(':nth-child(45) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(52) > .mantine-dj2qqh').contains('Name :')
        cy.get(':nth-child(53) > .mantine-dj2qqh').contains('Vicinity :')
        cy.get(':nth-child(54) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(58) > .mantine-dj2qqh').contains('Restaurant name :')
        cy.get(':nth-child(59) > .mantine-dj2qqh').contains('Address :')
        cy.get(':nth-child(60) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(64) > .mantine-dj2qqh').contains('Restaurant name :')
        cy.get(':nth-child(65) > .mantine-dj2qqh').contains('Address :')
        cy.get(':nth-child(66) > .mantine-dj2qqh').contains('Rating :')

        cy.get(':nth-child(70) > .mantine-dj2qqh').contains('Hotel :')
        cy.get(':nth-child(71) > .mantine-dj2qqh').contains('Address :')
        cy.get(':nth-child(72) > .mantine-dj2qqh').contains('Rating :')
    })
})
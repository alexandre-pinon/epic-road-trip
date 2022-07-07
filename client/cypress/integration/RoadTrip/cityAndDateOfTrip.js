describe('The cities and dates of the trip', () => {

    it('Check the start city is the same on the startEnd Page', () => {
        cy.visit('http://localhost:3000/')
        cy.fixture('tripInformations').then((tripInformations) => {
            //cy.log(tripInformations.startCity)
            cy.get('[data-testid="searchBar"]')
                .type(tripInformations.startCity+'{enter}')
                .type('{enter}')
            cy.wait(250)

            cy.location('href').should('include', '/startEndTrip')

            cy.get('[data-testid="StartCity"]').should('have.value', tripInformations.startCity)
        })
    })

    it('Check the text in the StarEndTrip page', () => {
        cy.visit('http://localhost:3000/startEndTrip')

        cy.get('[data-testid="title"]').contains('Start & End of your sub-trip')
        cy.get('[data-testid="cityTitle"]').contains('Choose the city of departure & arrival of your sub-trip')
        cy.get('[data-testid="dateTitle"]').contains('Choose the date of departure & arrival of your sub-trip')
    })

    it('Check the attribute in the StarEndTrip page', () => {
        cy.visit('http://localhost:3000/startEndTrip')

        cy.get('[data-testid="StartCity"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Start city of your Trip')

        cy.get('[data-testid="endCity"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'End city of your Trip')

        cy.get('[data-testid="startDate"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'Start date of your Trip')

        cy.get('[data-testid="endDate"]')
            .should('have.attr', 'placeholder')
            .and('equal', 'End date of your Trip')
    })



    it('Check the button in the StarEndTrip page', () => {
        cy.visit('http://localhost:3000/startEndTrip')

        cy.get('[data-testid="goBack"]').click()
        cy.wait(250)
        cy.location('href').should('include', '/')
    })



    it('Check the input in the StarEndTrip page', () => {
        cy.visit('http://localhost:3000/startEndTrip')

        cy.fixture('tripInformations').then((tripInformations) => {
            cy.get('[data-testid="StartCity"]')
                .type(tripInformations.startCity + '{enter}')
                .type('{enter}')

            cy.get('[data-testid="StartCity"]').should('have.value', tripInformations.startCity)

            cy.get('[data-testid="endCity"]')
                .type(tripInformations.endCity + '{enter}')
                .type('{enter}')

            cy.get('[data-testid="endCity"]').should('have.value', tripInformations.endCity)


            cy.get('[data-testid="startDate"]').click()
            cy.wait(250)
            cy.get(':nth-child(4) > :nth-child(3) > .mantine-388pmv').click()

            cy.get('[data-testid="startDate"]').should('have.value', 'July 20, 2022')


            cy.get('[data-testid="endDate"]').click()
            cy.wait(250)
            cy.get(':nth-child(5) > :nth-child(6) > .mantine-DatePicker-day').click()

            cy.get('[data-testid="endDate"]').should('have.value', 'July 30, 2022')
        })



    })









})
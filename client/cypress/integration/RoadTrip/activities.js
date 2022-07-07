describe('Activities pages', () => {

    it('Enjoy Activities', () => {
        cy.visit('http://localhost:3000/enjoy')

        cy.get('[data-testid="title"]').contains("Enjoy Activities")

        cy.get('[data-testid="goSleep"]').click()
        cy.location('href').should('include', '/sleep')
    })
})
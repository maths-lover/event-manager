package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomSponsorship(t *testing.T) Sponsorship {
	event := createRandomEvent(t)
	args := CreateSponsorshipParams{
		CompanyName: util.RandomString(10),
		Email:       util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Package: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		ContactPerson: util.RandomString(10),
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
	}

	sponsorship, err := testQueries.CreateSponsorship(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, sponsorship)

	require.NotZero(t, sponsorship.ID)
	require.NotZero(t, sponsorship.CompanyName)

	require.Equal(t, args.CompanyName, sponsorship.CompanyName)
	require.Equal(t, args.Email, sponsorship.Email)
	require.Equal(t, args.Phone, sponsorship.Phone)
	require.Equal(t, args.EventID, sponsorship.EventID)
	require.Equal(t, args.ContactPerson, sponsorship.ContactPerson)
	require.Equal(t, args.Package, sponsorship.Package)

	return sponsorship
}

func TestCreateSponsorship(t *testing.T) {
	createRandomSponsorship(t)
}

func TestListSponsorship(t *testing.T) {
	sponsorship1 := createRandomSponsorship(t)
	sponsorship2, err := testQueries.ListSponsorship(context.Background(), sponsorship1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, sponsorship2)

	require.Equal(t, reflect.DeepEqual(sponsorship2, sponsorship1), true)
}

func TestDeleteSponsorship(t *testing.T) {
	sponsorship1 := createRandomSponsorship(t)
	err := testQueries.DeleteSponsorship(context.Background(), sponsorship1.ID)
	require.NoError(t, err)

	sponsorship2, err := testQueries.ListSponsorship(context.Background(), sponsorship1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, sponsorship2)
}

func TestUpdateSponsorship(t *testing.T) {
	sponsorship1 := createRandomSponsorship(t)
	event := createRandomEvent(t)
	args := UpdateSponsorshipParams{
		ID:          sponsorship1.ID,
		CompanyName: util.RandomString(10),
		Email:       util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Package: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		ContactPerson: util.RandomString(10),
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
	}

	sponsorship2, err := testQueries.UpdateSponsorship(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, sponsorship2)

	require.Equal(t, args.CompanyName, sponsorship2.CompanyName)
	require.Equal(t, args.Email, sponsorship2.Email)
	require.Equal(t, args.Phone, sponsorship2.Phone)
	require.Equal(t, args.EventID, sponsorship2.EventID)
	require.Equal(t, args.ContactPerson, sponsorship2.ContactPerson)
	require.Equal(t, args.Package, sponsorship2.Package)
}

func TestListSponsorships(t *testing.T) {
	// create random events first
	for i := 0; i < 10; i++ {
		createRandomSponsorship(t)
	}

	args := ListSponsorshipsParams{
		Limit:  5,
		Offset: 5,
	}

	sponsorships, err := testQueries.ListSponsorships(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, sponsorships, 5)

	for _, sponsorship := range sponsorships {
		require.NotEmpty(t, sponsorship)
	}
}

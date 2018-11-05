package db

import (
	"github.com/devgit072/books-store/models"
	"github.com/golang/glog"
)

func GetBooks() ([]models.Book, error) {
	var book models.Book
	var books []models.Book
	db, err := ConnectDB()
	if err != nil {
		glog.Fatal(err)
		return nil, err
	}

	rows, err := db.Query("select id,title,author,publication,year from book")

	if err != nil {
		glog.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publication, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func GetBook(id int) (*models.Book, error) {
	db, err := ConnectDB()
	if err != nil {
		glog.Fatal(err)
		return nil, err
	}
    book := models.Book{}
	rows, err := db.Query("select id,title,author,publication,year from book where id=$1", id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Publication, &book.Year)
	} else {
		return nil, nil
	}
	return &book, nil
}

func AddBook(book *models.Book) (int,error) {
	db, err := ConnectDB()
	if err != nil {
		glog.Fatal(err)
		return -1,err
	}
	query := "INSERT INTO book(title,author,publication,year)VALUES($1,$2,$3,$4) RETURNING id"
	glog.Infof(query)
	var id int
	res := db.QueryRow(query, book.Title, book.Author, book.Publication, book.Year)
	err = res.Scan(&id)
	if err != nil {
		return -1,err
	}
	glog.Infof("Books inserted succesfully into table, with id: %d", id)
	return id, nil
}

func UpdateBook(book *models.Book) (int, error) {
	db,err := ConnectDB()
	if err != nil {
		glog.Fatal(err.Error())
		return -1, err
	}
	res, err := db.Exec("UPDATE book SET title=$1, author=$2, year=$3, publication=$4 where id=$5 RETURNING id",
		book.Title, book.Author, book.Year, book.Publication, book.ID)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		glog.Fatal(err.Error())
	}
	v := rowsAffected
	return int(v),nil
}

func RemoveBook(id int) (int, error) {
	db, err := ConnectDB()
	if err != nil {
		return -1, err
	}
	queryString := "DELETE FROM book WHERE id=$1"
	res, err := db.Exec(queryString, id)
	if err != nil {
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	return int(rowsAffected), err
}
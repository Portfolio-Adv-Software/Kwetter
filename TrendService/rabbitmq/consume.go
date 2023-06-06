package rabbitmq

import (
	pbtweet "github.com/Portfolio-Adv-Software/Kwetter/TrendService/proto"
	. "github.com/Portfolio-Adv-Software/Kwetter/TrendService/trendserver"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ConsumeMessage(queue string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := amqp.Dial("amqps://ctltdklj:qV9vx5HIf7JyfDDA0fRto3Disk-T57CF@goose.rmq2.cloudamqp.com/ctltdklj")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		for d := range msgs {
			tweet := &pbtweet.Tweet{}
			err := proto.Unmarshal(d.Body, tweet)
			if err != nil {
				log.Printf("failed to unmarshal tweet: %v", err)
				continue
			}
			tweet.Trend = extractHashtags(tweet.Body)
			log.Printf("received tweet: %v", tweet)
			client, _ := InitClient()
			PostTrend(client, tweet)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-c
}

func extractHashtags(tweetBody string) []string {
	// regular expression to match hashtags
	re := regexp.MustCompile(`#\w+`)

	// find all hashtags in the tweetContent
	hashtags := re.FindAllString(tweetBody, -1)

	uniqueHashtags := make(map[string]bool)

	//add each hashtag to a map
	for _, hashtag := range hashtags {
		// remove the "#" character from each hashtag
		hashtag = strings.TrimPrefix(hashtag, "#")

		//add hashtag to map if new
		if _, ok := uniqueHashtags[hashtag]; !ok {
			uniqueHashtags[hashtag] = true
		}
	}

	result := make([]string, 0, len(uniqueHashtags))
	for hashtag := range uniqueHashtags {
		result = append(result, hashtag)
	}
	return result
}

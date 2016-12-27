from rdflib import Graph
from time import time

g = Graph()

# Converting Tags.xml to Tags.nt

# Starting to parse Tags xml data from file.
print "Started parsing Tags.xml data from file."
start_time = time()
g.parse("data/Tags.xml", format='xml')
end_time = time()
print "End parsing Tags.xml data file. Time taken: %s (sec)" %(end_time-start_time)
print "Total ntriples length: %s" %(len(g))

# Starting to serialize Tags nquads data to file.
print "Started serializing Tags.nt data to file."
start_time = time()
g.serialize("data/Tags.nt", format='nt')
end_time = time()
print "End serializing Tags.nquads data file. Time taken: %s (sec)" %(end_time-start_time)


# Converting Users.xml to Users.nt

# Starting to parse Users xml data from file.
print "Started parsing Users.xml data from file."
start_time = time()
g.parse("data/Users.xml", format='xml')
end_time = time()
print "End parsing Users.xml data file. Time taken: %s (sec)" %(end_time-start_time)
print "Total ntriples length: %s" %(len(g))

# Starting to serialize Users nquads data to file.
print "Started serializing Users.nt data to file."
start_time = time()
g.serialize("data/Users.nt", format='nt')
end_time = time()
print "End serializing Users.nquads data file. Time taken: %s (sec)" %(end_time-start_time)


# Converting Posts.xml to Posts.nt

# Starting to parse Posts xml data from file.
print "Started parsing Posts.xml data from file."
start_time = time()
g.parse("data/Posts.xml", format='xml')
end_time = time()
print "End parsing Posts.xml data file. Time taken: %s (sec)" %(end_time-start_time)
print "Total ntriples length: %s" %(len(g))

# Starting to serialize Posts nquads data to file.
print "Started serializing Posts.nt data to file."
start_time = time()
g.serialize("data/Posts.nt", format='nt')
end_time = time()
print "End serializing Posts.nquads data file. Time taken: %s (sec)" %(end_time-start_time)
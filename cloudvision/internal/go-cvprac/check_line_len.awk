#!/usr/bin/awk -f
# Copyright (C) 2016  Arista Networks, Inc.

BEGIN {
  max = 100;
}

# Expand tabs to 4 spaces.
{
  gsub(/\t/, "    ");
}

length() > max {
  errors++;
  print FILENAME ":" FNR ": Line too long (" length() "/" max ")";
}

END {
  if (errors >= 125) {
    errors = 125;
  }
  exit errors;
}
